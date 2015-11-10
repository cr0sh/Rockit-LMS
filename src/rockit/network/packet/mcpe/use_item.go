package mcpe

import "bytes"

//UseItemPacket is a packet implements <TODO>
type UseItemPacket struct{}

//Encode encodes the packet
func (pk *UseItemPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk UseItemPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
