package packet

import (
	"bytes"
	"net"
	"rockit/util/logger"
)

//DataPacket will be used to process MCPE data packets, containing encapsulated packets.
//Buffer is separated from packet header. Should be appended manually.
type DataPacket struct {
	*bytes.Buffer
	SeqNumber           uint32
	Head                byte
	EncapsulatedPackets []EncapsulatedPacket
	Packets             [][]byte
}

//FromPacket returns 'decoded' data packet from given normal packet.
func FromPacket(pk Packet) (dp DataPacket, err error) {
	dp.Buffer = bytes.NewBuffer(pk.Bytes())
	dp.Head = pk.Head
	err = dp.Decode()
	return
}

//Encode encodes Packets slice and SeqNumber to raw buffer.
func (dp *DataPacket) Encode() Packet {
	dp.Buffer.WriteByte(dp.Head)
	PutLTriad(dp.SeqNumber, dp.Buffer)
	for _, pk := range dp.Packets {
		dp.Write(pk)
	}
	return Packet{Buffer: bytes.NewBuffer(dp.Bytes()[1:]), Head: dp.Bytes()[0], Address: *new(net.UDPAddr)}
}

//Decode decodes raw buffer to Packets slice and SeqNumber.
func (dp *DataPacket) Decode() (err error) {
	offset := 0
	maxlen := len(dp.Bytes())
	if dp.SeqNumber, err = ReadLTriad(dp.Buffer); err != nil {
		return Error{bytes.NewBuffer(append([]byte{dp.Head}, dp.Bytes()...)), err.Error()}
	}
	offset += 3
	for offset < maxlen {
		off := 0
		ep := new(EncapsulatedPacket)
		ep.Buffer = bytes.NewBuffer(dp.Bytes())
		var pk []byte
		if err = ep.Decapsulate(&off); err != nil {
			logger.Debug("Offset", off)
			return Error{bytes.NewBuffer(append([]byte{dp.Head}, dp.Bytes()...)), err.Error()}
		}
		pk = ep.Payload
		dp.Packets = append(dp.Packets, pk)
		dp.EncapsulatedPackets = append(dp.EncapsulatedPackets, *ep)
		offset += off
	}
	return nil
}

//TotalLen returns full length of packet
func (dp *DataPacket) TotalLen() (length int) {
	length = 4
	for _, pk := range dp.Packets {
		length += len(pk)
	}
	return
}
