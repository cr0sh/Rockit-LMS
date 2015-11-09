package mcpe

import "bytes"

//SetDifficultyPacket is a packet implements <TODO>
type SetDifficultyPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetDifficultyPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetDifficultyPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetDifficultyPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetDifficultyPacket) SetField(string) interface{} {
    return nil
}
