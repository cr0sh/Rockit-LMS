package mcpe

import "bytes"

//MoveEntityPacket is a packet implements <TODO>
type MoveEntityPacket struct{}

//Encode encodes the packet
func (pk *MoveEntityPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MoveEntityPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
