package mcpe

import "rockit/util/binary"

//ContainerSetContentPacket is a packet implements <TODO>
type ContainerSetContentPacket struct{}

//Encode encodes the packet
func (pk *ContainerSetContentPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerSetContentPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
