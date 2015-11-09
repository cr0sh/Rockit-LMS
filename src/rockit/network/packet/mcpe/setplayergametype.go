package mcpe

import "bytes"

//SetPlayerGametypePacket is a packet implements <TODO>
type SetPlayerGametypePacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetPlayerGametypePacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetPlayerGametypePacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetPlayerGametypePacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetPlayerGametypePacket) SetField(string) interface{} {
    return nil
}
