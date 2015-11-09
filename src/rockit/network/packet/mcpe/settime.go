package mcpe

import "bytes"

//SetTimePacket is a packet implements <TODO>
type SetTimePacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetTimePacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetTimePacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetTimePacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetTimePacket) SetField(string) interface{} {
    return nil
}
