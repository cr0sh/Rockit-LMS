package mcpe

import "bytes"

//TelemetryEventPacket is a packet implements <TODO>
type TelemetryEventPacket struct {
    *bytes.Buffer
    fields map[string]interface{}
}

//Encode encodes the packet
func (pk TelemetryEventPacket) Encode() error {
    return nil
}

//Decode decodes the packet
func (pk TelemetryEventPacket) Decode() error {
    return nil
}

//GetField returns specified field
func (pk TelemetryEventPacket) GetField(string) interface{} {
    return nil
}

//SetField sets specified field
func (pk TelemetryEventPacket) SetField(string) interface{} {
    return nil
}
