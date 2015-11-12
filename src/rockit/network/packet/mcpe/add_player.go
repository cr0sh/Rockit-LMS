package mcpe

import "rockit/util/binary"

//AddPlayerPacket is a packet implements <TODO>
type AddPlayerPacket struct{}

//Encode encodes the packet
func (pk *AddPlayerPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddPlayerPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
