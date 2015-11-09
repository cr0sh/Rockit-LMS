package mcpe

import "bytes"

//DropItemPacket is a packet implements <TODO>
type DropItemPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk DropItemPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk DropItemPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk DropItemPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk DropItemPacket) SetField(string) interface{} {
    return nil
}
