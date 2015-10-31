package packet

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"sort"
	"util/logging"
)

//Error is a implementation of error from packets.
type Error struct {
	buffer   *bytes.Buffer
	ErrorStr string
}

//Error implements error interface.
func (err Error) Error() string {
	return "Packet error(head 0x" + hex.EncodeToString([]byte{err.buffer.Bytes()[0]}) + "): " + err.ErrorStr + "\n" + hex.Dump(err.buffer.Bytes())
}

//NewError creates new packet error struct.
func NewError(buf *bytes.Buffer, err error) Error {
	return Error{buffer: buf, ErrorStr: err.Error()}
}

//ReadLTriad gets 3-byte LE Triad from buffer.
func ReadLTriad(buf *bytes.Buffer) (n uint32, err error) {
	b := buf.Next(3)
	if len(b) != 3 {
		err = fmt.Errorf("ReadLTriad: 3 bytes needed, %d given", len(b))
		return
	}
	n = uint32(b[0]) + uint32(b[1])<<8 + uint32(b[2])<<16
	return
}

//PutLTriad writes 3-byte LE Triad to buffer.
func PutLTriad(i uint32, buf *bytes.Buffer) (err error) {
	_, err = buf.Write([]byte{byte(i) & 0xff, byte(i >> 8), byte(i >> 16)})
	return
}

//PutAddress writes IP version, Address, Port from given net.UDPAddr struct to buffer. `version` is reserved for future IPv6 implementation.
func PutAddress(addr net.UDPAddr, buf *bytes.Buffer, version int) error {
	buf.WriteByte(4) // IPv4
	binary.Write(buf, binary.BigEndian, uint16(addr.Port))
	return nil
}

//ReadAddress reads IP version, Address, Port from given buffer.
func ReadAddress(buf *bytes.Buffer) (addr net.UDPAddr, err error) {
	var version byte
	if version, err = buf.ReadByte(); err != nil {
		return
	}
	if version == 4 {
		addrbuf := make([]byte, 4)
		if _, err = buf.Read(addrbuf); err != nil {
			return
		}
		for i := range addrbuf {
			addrbuf[i] ^= 0xff
		}
		addr.IP = addrbuf
		var port uint16
		if err = binary.Read(buf, binary.BigEndian, &port); err != nil {
			return
		}
		addr.Port = int(port)
		return
	}
	return addr, fmt.Errorf("IPv6 unsupported")
}

//Packet struct, with []bytes data and address(only for received packets)
//Buffer is separated from packet header. Should be appended manually.
type Packet struct {
	*bytes.Buffer
	Head    byte
	Address net.UDPAddr
}

//ReadStr reads string from packet.
func (p *Packet) ReadStr() (s string, err error) {
	var l uint16
	if err = binary.Read(p, binary.BigEndian, &l); err != nil {
		return
	}
	buf := make([]byte, l)
	if _, err = p.Read(buf); err != nil {
		return
	}
	s = string(buf)
	return
}

//PutStr converts string to MCPE-handlable bytes Buffer.
func (p *Packet) PutStr(s string) error {
	if len(s) > 65535 {
		return errors.New("String too long")
	}
	l := uint16(len(s))
	p.WriteByte(byte(l >> 8))
	p.WriteByte(byte(l & 0xff))
	_, err := p.Write([]byte(s))
	return err
}

//GetBytes returns bytes buffer from packet, with header.
func (p *Packet) GetBytes() []byte {
	return append([]byte{p.Head}, p.Buffer.Bytes()...)
}

//NewPacket returns empty packet with given header.
func NewPacket(head byte) Packet {
	return Packet{bytes.NewBuffer([]byte{}), head, *new(net.UDPAddr)}
}

//EncapsulatedPacket is raknet encapsulated packet, used with DataPacket.
type EncapsulatedPacket struct {
	*bytes.Buffer
	Reliability  byte
	HasSplit     bool
	Length       int16
	MessageIndex uint32
	OrderIndex   uint32
	OrderChannel byte
	SplitCount   uint32
	SplitID      uint16
	SplitIndex   uint32
	NeedACK      bool
}

//Encapsulate embeds packet to EncapsulatedPacket struct.
//Writes a packet to encapsulate, and options, and run this to get encapsulated packet buffer.
func (ep *EncapsulatedPacket) Encapsulate(p Packet) error {
	ep.Buffer = new(bytes.Buffer)
	flags := byte(ep.Reliability << 5)
	if ep.HasSplit {
		flags |= 1 << 4
	}
	ep.WriteByte(flags)
	if len(p.GetBytes()) >= 65536/8 {
		return fmt.Errorf("EncapsulatedPacket length field overflow")
	}
	binary.Write(ep.Buffer, binary.BigEndian, int16(len(p.GetBytes())<<3))
	if ep.Reliability > 0 {
		switch {
		case ep.Reliability >= 2 && ep.Reliability != 5:
			PutLTriad(ep.MessageIndex, ep.Buffer)
			fallthrough
		case ep.Reliability <= 4 && ep.Reliability != 2:
			PutLTriad(ep.OrderIndex, ep.Buffer)
			ep.WriteByte(ep.OrderChannel)
		}
	}
	if ep.HasSplit {
		binary.Write(ep.Buffer, binary.BigEndian, ep.SplitCount)
		binary.Write(ep.Buffer, binary.BigEndian, ep.SplitID)
		binary.Write(ep.Buffer, binary.BigEndian, ep.SplitIndex)
	}
	ep.Write(p.GetBytes())
	return nil
}

//Decapsulate extracts packet and gets options from EncapsulatedPacket buffer.
//Puts raw EncapsulatedPacket buffer to struct and run this to get decapsulated packet.
func (ep *EncapsulatedPacket) Decapsulate(offset *int) (pk Packet, err error) {
	pk = NewPacket(0)
	var flags byte
	if flags, err = ep.ReadByte(); err != nil {
		return pk, errors.New("Error while reading flags: " + err.Error())
	}
	*offset = 1
	ep.Reliability = (flags & (7 << 5)) >> 5
	ep.HasSplit = (flags & 16) > 0
	length := make([]byte, 2)
	var n int
	if n, err = ep.Read(length); n < 2 || err != nil {
		return pk, errors.New(err.Error())
	}
	*offset += 2
	if ep.Reliability > 0 {
		if ep.Reliability >= 2 && ep.Reliability != 5 {
			if ep.MessageIndex, err = ReadLTriad(ep.Buffer); err != nil {
				return pk, errors.New("Error while reading MessageIndex: " + err.Error())
			}
			*offset += 3
		}
		if ep.Reliability <= 4 && ep.Reliability != 2 {
			if ep.OrderIndex, err = ReadLTriad(ep.Buffer); err != nil {
				return pk, errors.New("Error while reading OrderIndex: " + err.Error())
			}
			*offset += 3
			if ep.OrderChannel, err = ep.ReadByte(); err != nil {
				return pk, errors.New("Error while reading OrderChannel: " + err.Error())
			}
			*offset++
		}
	}
	if ep.HasSplit {
		if err = binary.Read(ep.Buffer, binary.BigEndian, &ep.SplitCount); err != nil {
			return pk, errors.New("Error while reading SplitCount: " + err.Error())
		}
		*offset += 4
		if err = binary.Read(ep.Buffer, binary.BigEndian, &ep.SplitID); err != nil {
			return pk, errors.New("Error while reading SplitID: " + err.Error())
		}
		*offset += 2
		if err = binary.Read(ep.Buffer, binary.BigEndian, &ep.SplitIndex); err != nil {
			return pk, errors.New("Error while reading SplitIndex: " + err.Error())
		}
	}
	buf := make([]byte, binary.BigEndian.Uint16(length))
	if _, err = ep.Read(buf); err != nil {
		return pk, errors.New("Error while reading encapsulated buffer: " + err.Error())
	}
	*offset += int(binary.BigEndian.Uint16(length))
	if binary.BigEndian.Uint16(length) > 1 {
		pk.Buffer = bytes.NewBuffer(buf[1:])
	}
	pk.Head = buf[0]
	return
}

//Serializable specifies how to encode/decode packets to/from raw buffer.
type Serializable interface {
	Encode() error
	Decode() error
}

//DataPacket will be used to process MCPE data packets, containing encapsulated packets.
//Buffer is separated from packet header. Should be appended manually.
type DataPacket struct {
	*bytes.Buffer
	SeqNumber           uint32
	Head                byte
	EncapsulatedPackets []EncapsulatedPacket
	Packets             []Packet
}

//NewDataPacket returns 'decoded' data packet from given normal packet.
func NewDataPacket(pk Packet) (dp DataPacket, err error) {
	dp.Buffer = bytes.NewBuffer(pk.Bytes())
	dp.Head = pk.Head
	err = dp.Decode()
	return
}

//Encode encodes Packets slice and SeqNumber to raw buffer.
func (dp *DataPacket) Encode(head byte) Packet {
	dp.Buffer.WriteByte(head)
	PutLTriad(dp.SeqNumber, dp.Buffer)
	for _, pk := range dp.Packets {
		dp.Write(pk.Buffer.Bytes())
	}
	return Packet{Buffer: bytes.NewBuffer(dp.Bytes()[1:]), Head: dp.Bytes()[0], Address: *new(net.UDPAddr)}
}

//Decode decodes raw buffer to Packets slice and SeqNumber.
func (dp *DataPacket) Decode() (err error) {
	offset := 0
	if dp.SeqNumber, err = ReadLTriad(dp.Buffer); err != nil {
		return Error{bytes.NewBuffer(append([]byte{dp.Head}, dp.Bytes()...)), err.Error()}
	}
	offset += 3
	for offset < len(dp.Bytes()) {
		off := 0
		ep := new(EncapsulatedPacket)
		ep.Buffer = bytes.NewBuffer(dp.Bytes()[offset-3:])
		var pk Packet
		if pk, err = ep.Decapsulate(&off); err != nil {
			logging.Debug("Offset", off)
			return Error{bytes.NewBuffer(append([]byte{dp.Head}, dp.Bytes()...)), err.Error()}
		}
		dp.Packets = append(dp.Packets, pk)
		dp.EncapsulatedPackets = append(dp.EncapsulatedPackets, *ep)
		offset += off
	}
	return nil
}

type seqList []uint32

func (s seqList) Len() int {
	return len(s)
}

func (s seqList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s seqList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//AcknowledgePacket is a helper struct for encoding/decoding (N)ACK packets.
type AcknowledgePacket struct {
	*bytes.Buffer
	Packets seqList
}

//Encode encodes AcknowledgePacket.
func (a *AcknowledgePacket) Encode() {
	sort.Sort(a.Packets)
	payload := new(bytes.Buffer)
	count := a.Packets.Len()
	records := 0
	if count > 0 {
		pointer, start, last := 1, a.Packets[0], a.Packets[0]
		for pointer < count {
			current := a.Packets[pointer]
			pointer++
			diff := current - last
			if diff == 1 {
				last = current
			} else if diff > 1 {
				if start == last {
					payload.WriteByte(0x01)
					PutLTriad(start, payload)
					last = current
					start = last
				} else {
					payload.WriteByte(0x00)
					PutLTriad(start, payload)
					PutLTriad(start, payload)
					last = current
					start = last
				}
				records++
			}
		}
		if start == last {
			payload.WriteByte(0x01)
			PutLTriad(start, payload)
		} else {
			payload.WriteByte(0x00)
			PutLTriad(start, payload)
			PutLTriad(start, payload)
		}
		records++
	}
}

//Decode decodes AcknowledgePacket.
func (a *AcknowledgePacket) Decode() (err error) {
	var packetcount uint16
	if err = binary.Read(a.Buffer, binary.BigEndian, &packetcount); err != nil {
		return
	}
	cnt := 0
	a.Packets = make([]uint32, 0)
	var i uint16
	for ; i < packetcount && cnt < 4096; i++ {
		var flag byte
		if flag, err = a.ReadByte(); err != nil {
			return
		} else if flag == 0 {
			var start, end uint32
			if start, err = ReadLTriad(a.Buffer); err != nil {
				return
			}
			if end, err = ReadLTriad(a.Buffer); err != nil {
				return
			}
			if (end - start) > 512 {
				end = start + 512
			}
			for c := start; c <= end; c++ {
				a.Packets = append(a.Packets, c)
				cnt++
			}
		} else {
			var c uint32
			if c, err = ReadLTriad(a.Buffer); err != nil {
				return
			}
			a.Packets = append(a.Packets, c)
			cnt++
		}
	}
	return
}
