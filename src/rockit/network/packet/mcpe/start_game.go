package mcpe

import "bytes"

//StartGamePacket is a packet implements <TODO>
type StartGamePacket struct{}

//Encode encodes the packet
func (pk *StartGamePacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk StartGamePacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
