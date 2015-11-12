package mcpe

import "rockit/util/binary"

//PlayerInputPacket is a packet implements <TODO>
type PlayerInputPacket struct{}

//Encode encodes the packet
func (pk *PlayerInputPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk PlayerInputPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
