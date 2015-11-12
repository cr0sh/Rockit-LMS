package mcpe

import "rockit/util/binary"

//AnimatePacket is a packet implements <TODO>
type AnimatePacket struct{}

//Encode encodes the packet
func (pk *AnimatePacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AnimatePacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
