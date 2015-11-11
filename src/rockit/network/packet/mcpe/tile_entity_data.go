package mcpe

import "bytes"

//TileEntityDataPacket is a packet implements <TODO>
type TileEntityDataPacket struct{}

//Encode encodes the packet
func (pk *TileEntityDataPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TileEntityDataPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
