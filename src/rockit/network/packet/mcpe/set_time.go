package mcpe

import "bytes"

//SetTimePacket is a packet implements <TODO>
type SetTimePacket struct{}

//Encode encodes the packet
func (pk *SetTimePacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetTimePacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
