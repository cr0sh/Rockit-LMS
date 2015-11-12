package mcpe

import "rockit/util/binary"

//SetHealthPacket is a packet implements <TODO>
type SetHealthPacket struct{}

//Encode encodes the packet
func (pk *SetHealthPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetHealthPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
