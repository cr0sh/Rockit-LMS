package mcpe

import "bytes"

//SetSpawnPositionPacket is a packet implements <TODO>
type SetSpawnPositionPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk SetSpawnPositionPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk SetSpawnPositionPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk SetSpawnPositionPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk SetSpawnPositionPacket) SetField(string) interface{} {
    return nil
}
