package mcpe

import "rockit/util/binary"

//AddEntityPacket is a packet implements <TODO>
type AddEntityPacket struct{}

//Encode encodes the packet
func (pk *AddEntityPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AddEntityPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
