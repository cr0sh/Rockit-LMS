package mcpe

import "bytes"

//TextPacket is a packet implements <TODO>
type TextPacket struct{}

//Encode encodes the packet
func (pk *TextPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TextPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
