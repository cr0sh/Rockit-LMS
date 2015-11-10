package mcpe

import "bytes"

//TelemetryEventPacket is a packet implements <TODO>
type TelemetryEventPacket struct{}

//Encode encodes the packet
func (pk *TelemetryEventPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TelemetryEventPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
