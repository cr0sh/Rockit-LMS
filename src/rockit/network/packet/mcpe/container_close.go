package mcpe

import "rockit/util/binary"

//ContainerClosePacket is a packet implements <TODO>
type ContainerClosePacket struct{}

//Encode encodes the packet
func (pk *ContainerClosePacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerClosePacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
