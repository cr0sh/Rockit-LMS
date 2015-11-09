package mcpe

import "bytes"

//RemoveEntityPacket is a packet implements <TODO>
type RemoveEntityPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk RemoveEntityPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk RemoveEntityPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk RemoveEntityPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk RemoveEntityPacket) SetField(string) interface{} {
    return nil
}
