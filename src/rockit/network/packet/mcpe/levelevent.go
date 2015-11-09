package mcpe

import "bytes"

//LevelEventPacket is a packet implements <TODO>
type LevelEventPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk LevelEventPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk LevelEventPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk LevelEventPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk LevelEventPacket) SetField(string) interface{} {
    return nil
}
