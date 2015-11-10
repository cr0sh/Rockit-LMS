package mcpe

import "bytes"

//LoginPacket is a packet implements <TODO>
type LoginPacket struct{}

//Encode encodes the packet
func (pk *LoginPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk LoginPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
