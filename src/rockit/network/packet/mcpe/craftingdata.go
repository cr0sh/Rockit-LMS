package mcpe

import "bytes"

//CraftingDataPacket is a packet implements <TODO>
type CraftingDataPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk CraftingDataPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk CraftingDataPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk CraftingDataPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk CraftingDataPacket) SetField(string) interface{} {
    return nil
}
