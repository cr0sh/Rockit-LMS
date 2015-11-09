package mcpe

import "bytes"

//MobArmorEquipmentPacket is a packet implements <TODO>
type MobArmorEquipmentPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk MobArmorEquipmentPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk MobArmorEquipmentPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk MobArmorEquipmentPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk MobArmorEquipmentPacket) SetField(string) interface{} {
    return nil
}
