package mcpe

import "rockit/util/binary"

//DropItemPacket is a packet implements <TODO>
type DropItemPacket struct{}

//Encode encodes the packet
func (pk *DropItemPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk DropItemPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
