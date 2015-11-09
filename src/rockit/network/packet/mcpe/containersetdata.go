package mcpe

import "bytes"

//ContainerSetDataPacket is a packet implements <TODO>
type ContainerSetDataPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk ContainerSetDataPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk ContainerSetDataPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk ContainerSetDataPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk ContainerSetDataPacket) SetField(string) interface{} {
    return nil
}
