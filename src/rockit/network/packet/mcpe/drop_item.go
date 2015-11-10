package mcpe

import "bytes"

//DropItemPacket is a packet implements <TODO>
type DropItemPacket struct{}

//Encode encodes the packet
func (pk *DropItemPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk DropItemPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
