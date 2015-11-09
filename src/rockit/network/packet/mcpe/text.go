package mcpe

import "bytes"

//TextPacket is a packet implements <TODO>
type TextPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk TextPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk TextPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk TextPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk TextPacket) SetField(string) interface{} {
    return nil
}
