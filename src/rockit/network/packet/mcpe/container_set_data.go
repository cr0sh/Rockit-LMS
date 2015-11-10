package mcpe

import "bytes"

//ContainerSetDataPacket is a packet implements <TODO>
type ContainerSetDataPacket struct{}

//Encode encodes the packet
func (pk *ContainerSetDataPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerSetDataPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
