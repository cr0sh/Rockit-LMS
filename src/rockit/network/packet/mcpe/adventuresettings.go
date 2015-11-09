package mcpe

import "bytes"

//AdventureSettingsPacket is a packet implements <TODO>
type AdventureSettingsPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk AdventureSettingsPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk AdventureSettingsPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk AdventureSettingsPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk AdventureSettingsPacket) SetField(string) interface{} {
    return nil
}
