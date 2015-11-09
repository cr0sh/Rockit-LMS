package mcpe

import "bytes"

//RemoveBlockPacket is a packet implements <TODO>
type RemoveBlockPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk RemoveBlockPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk RemoveBlockPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk RemoveBlockPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk RemoveBlockPacket) SetField(string) interface{} {
    return nil
}
