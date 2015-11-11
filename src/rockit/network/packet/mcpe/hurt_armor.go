package mcpe

import "bytes"

//HurtArmorPacket is a packet implements <TODO>
type HurtArmorPacket struct{}

//Encode encodes the packet
func (pk *HurtArmorPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk HurtArmorPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
