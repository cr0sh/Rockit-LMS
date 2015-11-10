package mcpe

import "bytes"

//RemoveEntityPacket is a packet implements <TODO>
type RemoveEntityPacket struct{}

//Encode encodes the packet
func (pk *RemoveEntityPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RemoveEntityPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
