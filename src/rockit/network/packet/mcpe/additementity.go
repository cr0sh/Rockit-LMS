package mcpe

import "bytes"

//AddItemEntityPacket is a packet implements <TODO>
type AddItemEntityPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk AddItemEntityPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk AddItemEntityPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk AddItemEntityPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk AddItemEntityPacket) SetField(string) interface{} {
    return nil
}
