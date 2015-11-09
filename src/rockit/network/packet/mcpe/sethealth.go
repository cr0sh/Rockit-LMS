package mcpe

import "bytes"

//SetHealthPacket is a packet implements <TODO>
type SetHealthPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetHealthPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetHealthPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetHealthPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetHealthPacket) SetField(string) interface{} {
    return nil
}
