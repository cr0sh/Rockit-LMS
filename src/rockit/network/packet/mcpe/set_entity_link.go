package mcpe

import "rockit/util/binary"

//SetEntityLinkPacket is a packet implements <TODO>
type SetEntityLinkPacket struct{}

//Encode encodes the packet
func (pk *SetEntityLinkPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetEntityLinkPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
