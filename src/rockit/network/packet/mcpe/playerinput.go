package mcpe

import "bytes"

//PlayerInputPacket is a packet implements <TODO>
type PlayerInputPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk PlayerInputPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk PlayerInputPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk PlayerInputPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk PlayerInputPacket) SetField(string) interface{} {
    return nil
}
