package mcpe

import "bytes"

//BatchPacket is a packet implements <TODO>
type BatchPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk BatchPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk BatchPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk BatchPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk BatchPacket) SetField(string) interface{} {
    return nil
}
