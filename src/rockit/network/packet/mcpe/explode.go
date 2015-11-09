package mcpe

import "bytes"

//ExplodePacket is a packet implements <TODO>
type ExplodePacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk ExplodePacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk ExplodePacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk ExplodePacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk ExplodePacket) SetField(string) interface{} {
    return nil
}
