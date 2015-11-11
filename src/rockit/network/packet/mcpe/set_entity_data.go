package mcpe

import "bytes"

//SetEntityDataPacket is a packet implements <TODO>
type SetEntityDataPacket struct{}

//Encode encodes the packet
func (pk *SetEntityDataPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetEntityDataPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
