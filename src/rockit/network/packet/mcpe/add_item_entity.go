package mcpe

import "bytes"

//AddItemEntityPacket is a packet implements <TODO>
type AddItemEntityPacket struct{}

//Encode encodes the packet
func (pk *AddItemEntityPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddItemEntityPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
