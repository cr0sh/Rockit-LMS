package mcpe

import "bytes"

//EntityEventPacket is a packet implements <TODO>
type EntityEventPacket struct{}

//Encode encodes the packet
func (pk *EntityEventPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk EntityEventPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
