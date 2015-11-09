package mcpe

import "bytes"

//AddPaintingPacket is a packet implements <TODO>
type AddPaintingPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk AddPaintingPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk AddPaintingPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk AddPaintingPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk AddPaintingPacket) SetField(string) interface{} {
    return nil
}
