package mcpe

import "rockit/util/binary"

//TileEntityDataPacket is a packet implements <TODO>
type TileEntityDataPacket struct{}

//Encode encodes the packet
func (pk *TileEntityDataPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TileEntityDataPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
