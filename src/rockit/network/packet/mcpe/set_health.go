package mcpe

import "bytes"

//SetHealthPacket is a packet implements <TODO>
type SetHealthPacket struct{}

//Encode encodes the packet
func (pk *SetHealthPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetHealthPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
