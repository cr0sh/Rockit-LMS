package mcpe

import "rockit/util/binary"

//ContainerSetSlotPacket is a packet implements <TODO>
type ContainerSetSlotPacket struct{}

//Encode encodes the packet
func (pk *ContainerSetSlotPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerSetSlotPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
