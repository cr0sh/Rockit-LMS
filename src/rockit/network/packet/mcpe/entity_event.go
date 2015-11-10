package mcpe

import "bytes"

//EntityEventPacket is a packet implements <TODO>
type EntityEventPacket struct{}

//Encode encodes the packet
func (pk *EntityEventPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk EntityEventPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
