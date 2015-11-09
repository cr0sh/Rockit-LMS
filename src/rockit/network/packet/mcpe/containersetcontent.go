package mcpe

import "bytes"

//ContainerSetContentPacket is a packet implements <TODO>
type ContainerSetContentPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk ContainerSetContentPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk ContainerSetContentPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk ContainerSetContentPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk ContainerSetContentPacket) SetField(string) interface{} {
    return nil
}
