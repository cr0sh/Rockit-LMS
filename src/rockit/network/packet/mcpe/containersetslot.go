package mcpe

import "bytes"

//ContainerSetSlotPacket is a packet implements <TODO>
type ContainerSetSlotPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk ContainerSetSlotPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk ContainerSetSlotPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk ContainerSetSlotPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk ContainerSetSlotPacket) SetField(string) interface{} {
    return nil
}
