package mcpe

import "rockit/util/binary"

//CraftingEventPacket is a packet implements <TODO>
type CraftingEventPacket struct{}

//Encode encodes the packet
func (pk *CraftingEventPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk CraftingEventPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
