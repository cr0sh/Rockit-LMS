package mcpe

import "bytes"

//CraftingEventPacket is a packet implements <TODO>
type CraftingEventPacket struct{}

//Encode encodes the packet
func (pk *CraftingEventPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk CraftingEventPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
