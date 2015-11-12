package mcpe

import "rockit/util/binary"

//SetPlayerGametypePacket is a packet implements <TODO>
type SetPlayerGametypePacket struct{}

//Encode encodes the packet
func (pk *SetPlayerGametypePacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetPlayerGametypePacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
