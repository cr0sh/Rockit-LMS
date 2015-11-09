package mcpe

import "bytes"

//SetEntityLinkPacket is a packet implements <TODO>
type SetEntityLinkPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetEntityLinkPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetEntityLinkPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetEntityLinkPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetEntityLinkPacket) SetField(string) interface{} {
    return nil
}
