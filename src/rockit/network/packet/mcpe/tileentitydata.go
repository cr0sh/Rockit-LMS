package mcpe

import "bytes"

//TileEntityDataPacket is a packet implements <TODO>
type TileEntityDataPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk TileEntityDataPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk TileEntityDataPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk TileEntityDataPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk TileEntityDataPacket) SetField(string) interface{} {
    return nil
}
