package mcpe

import "bytes"

//ContainerClosePacket is a packet implements <TODO>
type ContainerClosePacket struct{}

//Encode encodes the packet
func (pk *ContainerClosePacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerClosePacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
