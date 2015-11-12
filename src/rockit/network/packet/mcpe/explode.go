package mcpe

import "rockit/util/binary"

//ExplodePacket is a packet implements <TODO>
type ExplodePacket struct{}

//Encode encodes the packet
func (pk *ExplodePacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk ExplodePacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
