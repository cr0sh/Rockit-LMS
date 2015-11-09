package mcpe

import "bytes"

//DisconnectPacket is a packet implements <TODO>
type DisconnectPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk DisconnectPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk DisconnectPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk DisconnectPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk DisconnectPacket) SetField(string) interface{} {
    return nil
}
