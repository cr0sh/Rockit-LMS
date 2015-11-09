package mcpe

import "bytes"

//PlayerActionPacket is a packet implements <TODO>
type PlayerActionPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk PlayerActionPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk PlayerActionPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk PlayerActionPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk PlayerActionPacket) SetField(string) interface{} {
    return nil
}
