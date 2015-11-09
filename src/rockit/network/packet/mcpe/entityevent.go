package mcpe

import "bytes"

//EntityEventPacket is a packet implements <TODO>
type EntityEventPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk EntityEventPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk EntityEventPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk EntityEventPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk EntityEventPacket) SetField(string) interface{} {
    return nil
}
