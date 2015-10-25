package packet

import (
	"bytes"
	"errors"
	"net"
)

//Packet struct, with []bytes data and address(only for received packets)
type Packet struct {
	*bytes.Buffer
	Head    byte
	Address net.UDPAddr
}

//EncapsulatedPacket is raknet encapsulated packet, used with DataPacket.
type EncapsulatedPacket struct {
	*bytes.Buffer
	Reliability  byte
	HasSplit     bool
	Length       int16
	MessageIndex []byte //3 bytes L-Triad
	OrderIndex   []byte //3 bytes L-Triad
	OrderChannel byte
	SplitCount   int32
	SplitID      int16
	SplitIndex   int32
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

//Bytes returns bytes buffer from packet, with header
func (p *Packet) Bytes() (b []byte) {
	b = []byte{p.Head}
	return append(b, p.Buffer.Bytes()...)
}

//New returns empty packet with given header
func New(head byte) Packet {
	return Packet{bytes.NewBuffer([]byte{}), head, *new(net.UDPAddr)}
}

//TODO: Implement Packet encapsulation/decapsulation

//Encapsulate returns encapsulated packet with given options
func Encapsulate(p Packet) EncapsulatedPacket {
	return *new(EncapsulatedPacket)
}

//Decapsulate returns decapsulated packet from given EncapsulatedPacket.
func Decapsulate(e EncapsulatedPacket) Packet {
	return *new(Packet)
}
