package mcpe

import "rockit/util/binary"

//MoveEntityPacket is a packet implements <TODO>
type MoveEntityPacket struct{}

//Encode encodes the packet
func (pk *MoveEntityPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MoveEntityPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
