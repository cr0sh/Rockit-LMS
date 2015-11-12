package mcpe

import "rockit/util/binary"

//RespawnPacket is a packet implements <TODO>
type RespawnPacket struct{}

//Encode encodes the packet
func (pk *RespawnPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RespawnPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
