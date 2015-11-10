package mcpe

import "bytes"

//TileEntityDataPacket is a packet implements <TODO>
type TileEntityDataPacket struct{}

//Encode encodes the packet
func (pk *TileEntityDataPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TileEntityDataPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
