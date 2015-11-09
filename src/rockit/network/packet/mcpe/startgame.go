package mcpe

import "bytes"

//StartGamePacket is a packet implements <TODO>
type StartGamePacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk StartGamePacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk StartGamePacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk StartGamePacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk StartGamePacket) SetField(string) interface{} {
    return nil
}
