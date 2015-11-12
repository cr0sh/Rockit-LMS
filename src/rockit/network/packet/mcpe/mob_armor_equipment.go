package mcpe

import "rockit/util/binary"

//MobArmorEquipmentPacket is a packet implements <TODO>
type MobArmorEquipmentPacket struct{}

//Encode encodes the packet
func (pk *MobArmorEquipmentPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MobArmorEquipmentPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
