package mcpe

import "bytes"

//TakeItemEntityPacket is a packet implements <TODO>
type TakeItemEntityPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk TakeItemEntityPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk TakeItemEntityPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk TakeItemEntityPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk TakeItemEntityPacket) SetField(string) interface{} {
    return nil
}
