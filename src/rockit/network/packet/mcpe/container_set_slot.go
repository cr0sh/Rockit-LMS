package mcpe

import "bytes"

//ContainerSetSlotPacket is a packet implements <TODO>
type ContainerSetSlotPacket struct{}

//Encode encodes the packet
func (pk *ContainerSetSlotPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerSetSlotPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
