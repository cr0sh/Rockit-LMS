package mcpe

import "bytes"

//UpdateAttributesPacket is a packet implements <TODO>
type UpdateAttributesPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk UpdateAttributesPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk UpdateAttributesPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk UpdateAttributesPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk UpdateAttributesPacket) SetField(string) interface{} {
    return nil
}
