package mcpe

import "rockit/util/binary"

//RemovePlayerPacket is a packet implements <TODO>
type RemovePlayerPacket struct{}

//Encode encodes the packet
func (pk *RemovePlayerPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RemovePlayerPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
