package mcpe

import "rockit/util/binary"

//SetSpawnPositionPacket is a packet implements <TODO>
type SetSpawnPositionPacket struct{}

//Encode encodes the packet
func (pk *SetSpawnPositionPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk SetSpawnPositionPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
