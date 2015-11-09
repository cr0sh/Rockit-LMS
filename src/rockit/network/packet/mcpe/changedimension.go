package mcpe

import "bytes"

//ChangeDimensionPacket is a packet implements <TODO>
type ChangeDimensionPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk ChangeDimensionPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk ChangeDimensionPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk ChangeDimensionPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk ChangeDimensionPacket) SetField(string) interface{} {
    return nil
}
