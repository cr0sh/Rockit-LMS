package mcpe

import "bytes"

//DisconnectPacket is a packet implements <TODO>
type DisconnectPacket struct{}

//Encode encodes the packet
func (pk *DisconnectPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk DisconnectPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
