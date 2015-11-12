package mcpe

import "rockit/util/binary"

//EntityEventPacket is a packet implements <TODO>
type EntityEventPacket struct{}

//Encode encodes the packet
func (pk *EntityEventPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk EntityEventPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
