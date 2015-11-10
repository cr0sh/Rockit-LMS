package mcpe

import "bytes"

//ContainerOpenPacket is a packet implements <TODO>
type ContainerOpenPacket struct{}

//Encode encodes the packet
func (pk *ContainerOpenPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ContainerOpenPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
