package mcpe

import "bytes"

//ExplodePacket is a packet implements <TODO>
type ExplodePacket struct{}

//Encode encodes the packet
func (pk *ExplodePacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ExplodePacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
