package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

//Packet struct, with []bytes data and address(only for received packets)
type Packet struct {
	*bytes.Buffer
	Head    byte
	Address net.UDPAddr
}

//WriteStr converts string to MCPE-handlable bytes Buffer
func (p *Packet) WriteStr(s string) error {
	if len(s) > 65535 {
		return errors.New("String too long")
	}
	l := uint16(len(s))
	p.WriteByte(byte(l >> 8))
	p.WriteByte(byte(l & 0xff))
	_, err := p.Write([]byte(s))
	return err
}

//GetBytes returns bytes buffer from packet, with header
func (p *Packet) GetBytes() []byte {
	return append([]byte{p.Head}, p.Buffer.Bytes()...)
}

//NewPacket returns empty packet with given header
func NewPacket(head byte) Packet {
	return Packet{bytes.NewBuffer([]byte{}), head, *new(net.UDPAddr)}
}

//ReadLTriad gets 3-byte LE Triad from buffer
func ReadLTriad(buf *bytes.Buffer) (n int32, err error) {
	b := buf.Next(3)
	if len(b) != 3 {
		err = fmt.Errorf("ReadLTriad: 3 bytes needed, %d given", len(b))
		return
	}
	n = int32(b[0]) + int32(b[2])<<8 + int32(b[3])<<16
	return
}

//PutLTriad writes 3-byte LE Triad to buffer
func PutLTriad(i int32, buf *bytes.Buffer) (err error) {
	_, err = buf.Write([]byte{byte(i) & 0xff, byte(i >> 8), byte(i >> 16)})
	return
}

//EncapsulatedPacket is raknet encapsulated packet, used with DataPacket.
type EncapsulatedPacket struct {
	*bytes.Buffer
	Reliability  byte
	HasSplit     bool
	Length       int16
	MessageIndex int32
	OrderIndex   int32
	OrderChannel byte
	SplitCount   int32
	SplitID      int16
	SplitIndex   int32
}

//TODO: Implement Packet encapsulation/decapsulation

//Encapsulate embeds packet to EncapsulatedPacket struct
func (ep *EncapsulatedPacket) Encapsulate(p Packet) error {
	ep.Buffer = new(bytes.Buffer)
	flags := ep.Reliability << 5
	if ep.HasSplit {
		flags |= 1 << 4
	}
	ep.Write([]byte{flags})
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

//Decapsulate extracts packet and gets options from EncapsulatedPacket buffer
func (ep *EncapsulatedPacket) Decapsulate(offset *int) (pk Packet, err error) {
	pk = NewPacket(0)
	*offset = 1
	var flags byte
	if flags, err = ep.ReadByte(); err != nil {
		return
	}
	ep.Reliability = (flags & (7 << 5)) >> 5
	ep.HasSplit = (flags & 16) > 0
	length := make([]byte, 2)
	*offset = 3
	var n int
	if n, err = ep.Read(length); n < 2 || err != nil {
		return
	}
	if ep.Reliability > 0 {
		switch {
		case ep.Reliability >= 2 && ep.Reliability != 5:
			if ep.MessageIndex, err = ReadLTriad(ep.Buffer); err != nil {
				return
			}
			*offset += 3
			fallthrough
		case ep.Reliability <= 4 && ep.Reliability != 2:
			if ep.OrderIndex, err = ReadLTriad(ep.Buffer); err != nil {
				return
			}
			*offset += 3
			if ep.OrderChannel, err = ep.ReadByte(); err != nil {
				return
			}
			*offset++
		}
	}
	if ep.HasSplit {
		if err = binary.Read(ep.Buffer, binary.BigEndian, ep.SplitCount); err != nil {
			return
		}
		*offset += 4
		if err = binary.Read(ep.Buffer, binary.BigEndian, ep.SplitID); err != nil {
			return
		}
		*offset += 2
		if err = binary.Read(ep.Buffer, binary.BigEndian, ep.SplitIndex); err != nil {
			return
		}
	}
	buf := make([]byte, binary.BigEndian.Uint16(length))
	if n, err = ep.Read(buf); n < int(binary.BigEndian.Uint16(length)) || err != nil {
		if err != nil {
			return
		}
		return pk, io.EOF
	}
	*offset += int(binary.BigEndian.Uint16(length))
	pk.Buffer = bytes.NewBuffer(buf[1:])
	pk.Head = buf[0]
	return
}
