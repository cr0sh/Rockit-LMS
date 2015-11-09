package mcpe

import "bytes"

//UseItemPacket is a packet implements <TODO>
type UseItemPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk UseItemPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk UseItemPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk UseItemPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk UseItemPacket) SetField(string) interface{} {
    return nil
}
