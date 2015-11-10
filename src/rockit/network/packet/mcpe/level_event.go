package mcpe

import "bytes"

//LevelEventPacket is a packet implements <TODO>
type LevelEventPacket struct{}

//Encode encodes the packet
func (pk *LevelEventPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk LevelEventPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
