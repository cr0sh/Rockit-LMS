package mcpe

import "bytes"

//SetSpawnPositionPacket is a packet implements <TODO>
type SetSpawnPositionPacket struct{}

//Encode encodes the packet
func (pk *SetSpawnPositionPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetSpawnPositionPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
