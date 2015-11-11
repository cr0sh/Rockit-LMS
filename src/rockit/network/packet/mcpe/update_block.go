package mcpe

import "bytes"

//UpdateBlockPacket is a packet implements <TODO>
type UpdateBlockPacket struct{}

//Encode encodes the packet
func (pk *UpdateBlockPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk UpdateBlockPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
