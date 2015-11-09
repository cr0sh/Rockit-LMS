package mcpe

import "bytes"

//MobEquipmentPacket is a packet implements <TODO>
type MobEquipmentPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk MobEquipmentPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk MobEquipmentPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk MobEquipmentPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk MobEquipmentPacket) SetField(string) interface{} {
    return nil
}
