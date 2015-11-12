package mcpe

import "rockit/util/binary"

//UseItemPacket is a packet implements <TODO>
type UseItemPacket struct{}

//Encode encodes the packet
func (pk *UseItemPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk UseItemPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
