package mcpe

import "bytes"

//AddItemEntityPacket is a packet implements <TODO>
type AddItemEntityPacket struct{}

//Encode encodes the packet
func (pk *AddItemEntityPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddItemEntityPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
