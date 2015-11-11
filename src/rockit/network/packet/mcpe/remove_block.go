package mcpe

import "bytes"

//RemoveBlockPacket is a packet implements <TODO>
type RemoveBlockPacket struct{}

//Encode encodes the packet
func (pk *RemoveBlockPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RemoveBlockPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
