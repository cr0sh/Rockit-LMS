package mcpe

import "bytes"

//PlayerActionPacket is a packet implements <TODO>
type PlayerActionPacket struct{}

//Encode encodes the packet
func (pk *PlayerActionPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk PlayerActionPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
