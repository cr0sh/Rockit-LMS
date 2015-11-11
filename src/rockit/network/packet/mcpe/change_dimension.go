package mcpe

import "bytes"

//ChangeDimensionPacket is a packet implements <TODO>
type ChangeDimensionPacket struct{}

//Encode encodes the packet
func (pk *ChangeDimensionPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ChangeDimensionPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
