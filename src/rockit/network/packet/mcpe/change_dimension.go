package mcpe

import "bytes"

//ChangeDimensionPacket is a packet implements <TODO>
type ChangeDimensionPacket struct{}

//Encode encodes the packet
func (pk *ChangeDimensionPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ChangeDimensionPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
