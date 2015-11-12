package mcpe

import "rockit/util/binary"

//SetDifficultyPacket is a packet implements <TODO>
type SetDifficultyPacket struct{}

//Encode encodes the packet
func (pk *SetDifficultyPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetDifficultyPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
