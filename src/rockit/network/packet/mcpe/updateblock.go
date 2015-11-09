package mcpe

import "bytes"

//UpdateBlockPacket is a packet implements <TODO>
type UpdateBlockPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk UpdateBlockPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk UpdateBlockPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk UpdateBlockPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk UpdateBlockPacket) SetField(string) interface{} {
    return nil
}
