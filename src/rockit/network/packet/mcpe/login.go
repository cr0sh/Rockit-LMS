package mcpe

import "bytes"

//LoginPacket is a packet implements <TODO>
type LoginPacket struct{}

//Encode encodes the packet
func (pk *LoginPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk LoginPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	fields = make(Field)
	return
}
