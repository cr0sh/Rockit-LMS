package mcpe

import "bytes"

//CraftingEventPacket is a packet implements <TODO>
type CraftingEventPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk CraftingEventPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk CraftingEventPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk CraftingEventPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk CraftingEventPacket) SetField(string) interface{} {
    return nil
}
