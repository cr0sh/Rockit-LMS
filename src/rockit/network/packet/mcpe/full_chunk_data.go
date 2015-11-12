package mcpe

import "rockit/util/binary"

//FullChunkDataPacket is a packet implements <TODO>
type FullChunkDataPacket struct{}

//Encode encodes the packet
func (pk *FullChunkDataPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk FullChunkDataPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
