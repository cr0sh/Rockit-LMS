package mcpe

import "rockit/util/binary"

//SetTimePacket is a packet implements <TODO>
type SetTimePacket struct{}

//Encode encodes the packet
func (pk *SetTimePacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetTimePacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
