package mcpe

import "rockit/util/binary"

//LevelEventPacket is a packet implements <TODO>
type LevelEventPacket struct{}

//Encode encodes the packet
func (pk *LevelEventPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk LevelEventPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
