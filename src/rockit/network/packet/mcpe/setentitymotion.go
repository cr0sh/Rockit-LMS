package mcpe

import "bytes"

//SetEntityMotionPacket is a packet implements <TODO>
type SetEntityMotionPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetEntityMotionPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetEntityMotionPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetEntityMotionPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetEntityMotionPacket) SetField(string) interface{} {
    return nil
}
