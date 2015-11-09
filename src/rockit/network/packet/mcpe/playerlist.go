package mcpe

import "bytes"

//PlayerListPacket is a packet implements <TODO>
type PlayerListPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk PlayerListPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk PlayerListPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk PlayerListPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk PlayerListPacket) SetField(string) interface{} {
    return nil
}
