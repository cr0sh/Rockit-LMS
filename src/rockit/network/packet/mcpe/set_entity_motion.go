package mcpe

import "rockit/util/binary"

//SetEntityMotionPacket is a packet implements <TODO>
type SetEntityMotionPacket struct{}

//Encode encodes the packet
func (pk *SetEntityMotionPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetEntityMotionPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
