package mcpe

import "bytes"

//TakeItemEntityPacket is a packet implements <TODO>
type TakeItemEntityPacket struct{}

//Encode encodes the packet
func (pk *TakeItemEntityPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TakeItemEntityPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}