package mcpe

import "bytes"

//AddEntityPacket is a packet implements <TODO>
type AddEntityPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk AddEntityPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk AddEntityPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk AddEntityPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk AddEntityPacket) SetField(string) interface{} {
    return nil
}
