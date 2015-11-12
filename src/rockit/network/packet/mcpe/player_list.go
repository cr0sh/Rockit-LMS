package mcpe

import "rockit/util/binary"

//PlayerListPacket is a packet implements <TODO>
type PlayerListPacket struct{}

//Encode encodes the packet
func (pk *PlayerListPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk PlayerListPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
