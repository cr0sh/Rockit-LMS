//Package packet defines packet structs, related functions for Raknet connection
package packet

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
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
	_, err = buf.Write([]byte{byte(i >> 16), byte(i >> 8), byte(i)})
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
