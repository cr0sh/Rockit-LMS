package mcpe

import "bytes"

//PlayerListPacket is a packet implements <TODO>
type PlayerListPacket struct{}

//Encode encodes the packet
func (pk *PlayerListPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk PlayerListPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
