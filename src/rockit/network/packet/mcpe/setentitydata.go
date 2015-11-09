package mcpe

import "bytes"

//SetEntityDataPacket is a packet implements <TODO>
type SetEntityDataPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetEntityDataPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetEntityDataPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetEntityDataPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetEntityDataPacket) SetField(string) interface{} {
    return nil
}
