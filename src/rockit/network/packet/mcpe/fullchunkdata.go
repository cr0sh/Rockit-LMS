package mcpe

import "bytes"

//FullChunkDataPacket is a packet implements <TODO>
type FullChunkDataPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk FullChunkDataPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk FullChunkDataPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk FullChunkDataPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk FullChunkDataPacket) SetField(string) interface{} {
    return nil
}
