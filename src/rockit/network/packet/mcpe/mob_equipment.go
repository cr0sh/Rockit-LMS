package mcpe

import "rockit/util/binary"

//MobEquipmentPacket is a packet implements <TODO>
type MobEquipmentPacket struct{}

//Encode encodes the packet
func (pk *MobEquipmentPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MobEquipmentPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
