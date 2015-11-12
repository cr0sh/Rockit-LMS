package mcpe

import "rockit/util/binary"

//ContainerSetDataPacket is a packet implements <TODO>
type ContainerSetDataPacket struct{}

//Encode encodes the packet
func (pk *ContainerSetDataPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerSetDataPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
