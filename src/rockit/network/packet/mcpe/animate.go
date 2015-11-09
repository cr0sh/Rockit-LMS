package mcpe

import "bytes"

//AnimatePacket is a packet implements <TODO>
type AnimatePacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk AnimatePacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk AnimatePacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk AnimatePacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk AnimatePacket) SetField(string) interface{} {
    return nil
}
