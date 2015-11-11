package mcpe

import "bytes"

//ContainerClosePacket is a packet implements <TODO>
type ContainerClosePacket struct{}

//Encode encodes the packet
func (pk *ContainerClosePacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerClosePacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
