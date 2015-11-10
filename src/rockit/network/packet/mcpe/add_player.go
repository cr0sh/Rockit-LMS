package mcpe

import "bytes"

//AddPlayerPacket is a packet implements <TODO>
type AddPlayerPacket struct{}

//Encode encodes the packet
func (pk *AddPlayerPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddPlayerPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
