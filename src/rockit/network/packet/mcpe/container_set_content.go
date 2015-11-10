package mcpe

import "bytes"

//ContainerSetContentPacket is a packet implements <TODO>
type ContainerSetContentPacket struct{}

//Encode encodes the packet
func (pk *ContainerSetContentPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerSetContentPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
