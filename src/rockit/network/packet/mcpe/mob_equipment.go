package mcpe

import "bytes"

//MobEquipmentPacket is a packet implements <TODO>
type MobEquipmentPacket struct{}

//Encode encodes the packet
func (pk *MobEquipmentPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk MobEquipmentPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
