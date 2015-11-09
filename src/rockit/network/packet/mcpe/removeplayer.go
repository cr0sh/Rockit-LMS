package mcpe

import "bytes"

//RemovePlayerPacket is a packet implements <TODO>
type RemovePlayerPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk RemovePlayerPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk RemovePlayerPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk RemovePlayerPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk RemovePlayerPacket) SetField(string) interface{} {
    return nil
}
