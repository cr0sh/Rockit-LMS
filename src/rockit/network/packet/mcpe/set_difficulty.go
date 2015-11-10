package mcpe

import "bytes"

//SetDifficultyPacket is a packet implements <TODO>
type SetDifficultyPacket struct{}

//Encode encodes the packet
func (pk *SetDifficultyPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetDifficultyPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
