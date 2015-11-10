package mcpe

import "bytes"

//RespawnPacket is a packet implements <TODO>
type RespawnPacket struct{}

//Encode encodes the packet
func (pk *RespawnPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RespawnPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
