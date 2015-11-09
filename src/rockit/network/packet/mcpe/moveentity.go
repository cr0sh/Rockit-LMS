package mcpe

import "bytes"

//MoveEntityPacket is a packet implements <TODO>
type MoveEntityPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk MoveEntityPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk MoveEntityPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk MoveEntityPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk MoveEntityPacket) SetField(string) interface{} {
    return nil
}
