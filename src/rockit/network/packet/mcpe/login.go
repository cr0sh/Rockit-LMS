package mcpe

import "bytes"

//LoginPacket is a packet implements <TODO>
type LoginPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk LoginPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk LoginPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk LoginPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk LoginPacket) SetField(string) interface{} {
    return nil
}
