package mcpe

import "bytes"

//RemovePlayerPacket is a packet implements <TODO>
type RemovePlayerPacket struct{}

//Encode encodes the packet
func (pk *RemovePlayerPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RemovePlayerPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
