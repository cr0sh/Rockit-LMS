package mcpe

import "bytes"

//InteractPacket is a packet implements <TODO>
type InteractPacket struct{}

//Encode encodes the packet
func (pk *InteractPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk InteractPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
