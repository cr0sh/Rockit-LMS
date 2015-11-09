package mcpe

import "bytes"

//InteractPacket is a packet implements <TODO>
type InteractPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk InteractPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk InteractPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk InteractPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk InteractPacket) SetField(string) interface{} {
    return nil
}
