package mcpe

import "bytes"

//TileEventPacket is a packet implements <TODO>
type TileEventPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk TileEventPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk TileEventPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk TileEventPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk TileEventPacket) SetField(string) interface{} {
    return nil
}
