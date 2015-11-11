package mcpe

import "bytes"

//PlayStatusPacket is a packet implements <TODO>
type PlayStatusPacket struct{}

//Encode encodes the packet
func (pk *PlayStatusPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk PlayStatusPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
