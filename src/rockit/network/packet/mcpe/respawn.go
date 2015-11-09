package mcpe

import "bytes"

//RespawnPacket is a packet implements <TODO>
type RespawnPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk RespawnPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk RespawnPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk RespawnPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk RespawnPacket) SetField(string) interface{} {
    return nil
}
