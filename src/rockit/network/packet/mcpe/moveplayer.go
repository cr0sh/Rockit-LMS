package mcpe

import "bytes"

//MovePlayerPacket is a packet implements <TODO>
type MovePlayerPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk MovePlayerPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk MovePlayerPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk MovePlayerPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk MovePlayerPacket) SetField(string) interface{} {
    return nil
}
