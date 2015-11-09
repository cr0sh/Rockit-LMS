package mcpe

import "bytes"

//HurtArmorPacket is a packet implements <TODO>
type HurtArmorPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk HurtArmorPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk HurtArmorPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk HurtArmorPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk HurtArmorPacket) SetField(string) interface{} {
    return nil
}
