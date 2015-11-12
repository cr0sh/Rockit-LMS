package mcpe

import "rockit/util/binary"

//MovePlayerPacket is a packet implements <TODO>
type MovePlayerPacket struct{}

//Encode encodes the packet
func (pk *MovePlayerPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MovePlayerPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
