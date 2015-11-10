package mcpe

import "bytes"

//MovePlayerPacket is a packet implements <TODO>
type MovePlayerPacket struct{}

//Encode encodes the packet
func (pk *MovePlayerPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MovePlayerPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
