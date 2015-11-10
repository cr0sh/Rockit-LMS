package mcpe

import "bytes"

//PlayerActionPacket is a packet implements <TODO>
type PlayerActionPacket struct{}

//Encode encodes the packet
func (pk *PlayerActionPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk PlayerActionPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
