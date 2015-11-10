package mcpe

import "bytes"

//AddPaintingPacket is a packet implements <TODO>
type AddPaintingPacket struct{}

//Encode encodes the packet
func (pk *AddPaintingPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddPaintingPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
