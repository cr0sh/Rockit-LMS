package mcpe

import "bytes"

//AddEntityPacket is a packet implements <TODO>
type AddEntityPacket struct{}

//Encode encodes the packet
func (pk *AddEntityPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddEntityPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
