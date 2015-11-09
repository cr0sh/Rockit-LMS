package mcpe

import "bytes"

//AddPlayerPacket is a packet implements <TODO>
type AddPlayerPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk AddPlayerPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk AddPlayerPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk AddPlayerPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk AddPlayerPacket) SetField(string) interface{} {
    return nil
}
