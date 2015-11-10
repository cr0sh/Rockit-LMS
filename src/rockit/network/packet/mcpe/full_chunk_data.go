package mcpe

import "bytes"

//FullChunkDataPacket is a packet implements <TODO>
type FullChunkDataPacket struct{}

//Encode encodes the packet
func (pk *FullChunkDataPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk FullChunkDataPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
