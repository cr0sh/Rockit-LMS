package mcpe

import "bytes"

//PlayStatusPacket is a packet implements <TODO>
type PlayStatusPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk PlayStatusPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk PlayStatusPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk PlayStatusPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk PlayStatusPacket) SetField(string) interface{} {
    return nil
}
