package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

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
	Payload      []byte
	Locked       bool // Do not modify
}

//Encapsulate embeds payload to EncapsulatedPacket struct.
//Writes a packet to encapsulate, and options, and run this to get encapsulated packet buffer.
func (ep *EncapsulatedPacket) Encapsulate() error {
	if ep.Locked {
		return nil
	}
	ep.Buffer = new(bytes.Buffer)
	flags := byte(ep.Reliability << 5)
	if ep.HasSplit {
		flags |= 1 << 4
	}
	ep.WriteByte(flags)
	if len(ep.Payload) >= 65536/8 {
		return fmt.Errorf("EncapsulatedPacket length field overflow")
	}
	binary.Write(ep.Buffer, binary.BigEndian, int16(len(ep.Payload)<<3))
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
	ep.Write(ep.Payload)
	return nil
}

//Decapsulate extracts payload and gets options from EncapsulatedPacket buffer.
//Puts raw EncapsulatedPacket buffer to struct and run this to get decapsulated packet.
func (ep *EncapsulatedPacket) Decapsulate(offset *int) (err error) {
	if ep.Locked {
		return
	}
	cap := ep.Len()
	var flags byte
	if flags, err = ep.ReadByte(); err != nil {
		return errors.New("Error while reading flags: " + err.Error())
	}
	*offset = 1
	ep.Reliability = (flags & (7 << 5)) >> 5
	ep.HasSplit = (flags & 16) > 0
	length := make([]byte, 2)
	var n int
	if n, err = ep.Read(length); n < 2 || err != nil {
		return errors.New(err.Error())
	}
	*offset += 2
	if ep.Reliability > 0 {
		if ep.Reliability >= 2 && ep.Reliability != 5 {
			if ep.MessageIndex, err = ReadLTriad(ep.Buffer); err != nil {
				return errors.New("Error while reading MessageIndex: " + err.Error())
			}
			*offset += 3
		}
		if ep.Reliability <= 4 && ep.Reliability != 2 {
			if ep.OrderIndex, err = ReadLTriad(ep.Buffer); err != nil {
				return errors.New("Error while reading OrderIndex: " + err.Error())
			}
			*offset += 3
			if ep.OrderChannel, err = ep.ReadByte(); err != nil {
				return errors.New("Error while reading OrderChannel: " + err.Error())
			}
			*offset++
		}
	}
	if ep.HasSplit {
		if err = binary.Read(ep.Buffer, binary.BigEndian, &ep.SplitCount); err != nil {
			return errors.New("Error while reading SplitCount: " + err.Error())
		}
		*offset += 4
		if err = binary.Read(ep.Buffer, binary.BigEndian, &ep.SplitID); err != nil {
			return errors.New("Error while reading SplitID: " + err.Error())
		}
		*offset += 2
		if err = binary.Read(ep.Buffer, binary.BigEndian, &ep.SplitIndex); err != nil {
			return errors.New("Error while reading SplitIndex: " + err.Error())
		}
		*offset += 4
	}
	if binary.BigEndian.Uint16(length) > uint16(cap-*offset) {
		ep.Payload = make([]byte, cap-*offset)
	} else {
		ep.Payload = make([]byte, binary.BigEndian.Uint16(length))
	}
	if _, err = ep.Read(ep.Payload); err != nil {
		return errors.New("Error while reading encapsulated buffer: " + err.Error())
	}
	*offset += int(binary.BigEndian.Uint16(length))
	return
}

//TotalLen returns total packet length
func (ep *EncapsulatedPacket) TotalLen() (len int) {
	len = 3
	len += ep.Len()
	if ep.Reliability >= 2 && ep.Reliability != 5 {
		len += 3
	}
	if ep.Reliability <= 4 && ep.Reliability != 2 {
		len += 4
	}
	if ep.HasSplit {
		len += 10
	}
	return
}
