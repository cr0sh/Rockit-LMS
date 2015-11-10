package mcpe

import "bytes"

//MobArmorEquipmentPacket is a packet implements <TODO>
type MobArmorEquipmentPacket struct{}

//Encode encodes the packet
func (pk *MobArmorEquipmentPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MobArmorEquipmentPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
