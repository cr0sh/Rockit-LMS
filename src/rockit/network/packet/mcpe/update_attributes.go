package mcpe

import "rockit/util/binary"

//UpdateAttributesPacket is a packet implements <TODO>
type UpdateAttributesPacket struct{}

//Encode encodes the packet
func (pk *UpdateAttributesPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk UpdateAttributesPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
