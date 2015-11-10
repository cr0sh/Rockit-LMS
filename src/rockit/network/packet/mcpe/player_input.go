package mcpe

import "bytes"

//PlayerInputPacket is a packet implements <TODO>
type PlayerInputPacket struct{}

//Encode encodes the packet
func (pk *PlayerInputPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk PlayerInputPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
