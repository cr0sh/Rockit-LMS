package mcpe

import "bytes"

//ContainerClosePacket is a packet implements <TODO>
type ContainerClosePacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk ContainerClosePacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk ContainerClosePacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk ContainerClosePacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk ContainerClosePacket) SetField(string) interface{} {
    return nil
}
