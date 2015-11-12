package mcpe

import "rockit/util/binary"

//ContainerOpenPacket is a packet implements <TODO>
type ContainerOpenPacket struct{}

//Encode encodes the packet
func (pk *ContainerOpenPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerOpenPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
