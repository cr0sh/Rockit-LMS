package mcpe

import "rockit/util/binary"

//StartGamePacket is a packet implements <TODO>
type StartGamePacket struct{}

//Encode encodes the packet
func (pk *StartGamePacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk StartGamePacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
