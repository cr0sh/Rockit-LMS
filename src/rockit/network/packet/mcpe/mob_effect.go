package mcpe

import "bytes"

//MobEffectPacket is a packet implements <TODO>
type MobEffectPacket struct{}

//Encode encodes the packet
func (pk *MobEffectPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MobEffectPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
