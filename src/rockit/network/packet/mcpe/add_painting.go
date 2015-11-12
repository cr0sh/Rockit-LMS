package mcpe

import "rockit/util/binary"

//AddPaintingPacket is a packet implements <TODO>
type AddPaintingPacket struct{}

//Encode encodes the packet
func (pk *AddPaintingPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddPaintingPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
