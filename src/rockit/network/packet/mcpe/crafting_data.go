package mcpe

import "bytes"

//CraftingDataPacket is a packet implements <TODO>
type CraftingDataPacket struct{}

//Encode encodes the packet
func (pk *CraftingDataPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk CraftingDataPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
