package mcpe

import "bytes"

//MobEffectPacket is a packet implements <TODO>
type MobEffectPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk MobEffectPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk MobEffectPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk MobEffectPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk MobEffectPacket) SetField(string) interface{} {
    return nil
}
