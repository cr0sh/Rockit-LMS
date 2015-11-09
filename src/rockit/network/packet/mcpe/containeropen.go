package mcpe

import "bytes"

//ContainerOpenPacket is a packet implements <TODO>
type ContainerOpenPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk ContainerOpenPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk ContainerOpenPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk ContainerOpenPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk ContainerOpenPacket) SetField(string) interface{} {
    return nil
}
