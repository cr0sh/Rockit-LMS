package mcpe

import "bytes"

//TileEventPacket is a packet implements <TODO>
type TileEventPacket struct{}

//Encode encodes the packet
func (pk *TileEventPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TileEventPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
