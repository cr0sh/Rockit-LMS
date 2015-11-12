package mcpe

import "rockit/util/binary"

//MobEffectPacket is a packet implements <TODO>
type MobEffectPacket struct{}

//Encode encodes the packet
func (pk *MobEffectPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MobEffectPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
