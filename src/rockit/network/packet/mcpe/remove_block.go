package mcpe

import "bytes"

//RemoveBlockPacket is a packet implements <TODO>
type RemoveBlockPacket struct{}

//Encode encodes the packet
func (pk *RemoveBlockPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RemoveBlockPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
