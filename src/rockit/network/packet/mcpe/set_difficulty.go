package mcpe

import "bytes"

//SetDifficultyPacket is a packet implements <TODO>
type SetDifficultyPacket struct{}

//Encode encodes the packet
func (pk *SetDifficultyPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetDifficultyPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
