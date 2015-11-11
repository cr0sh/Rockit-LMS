package mcpe

import "bytes"

//TelemetryEventPacket is a packet implements <TODO>
type TelemetryEventPacket struct{}

//Encode encodes the packet
func (pk *TelemetryEventPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk TelemetryEventPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
