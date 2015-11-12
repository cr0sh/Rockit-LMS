package mcpe

import "rockit/util/binary"

//RemoveBlockPacket is a packet implements <TODO>
type RemoveBlockPacket struct{}

//Encode encodes the packet
func (pk *RemoveBlockPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RemoveBlockPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
