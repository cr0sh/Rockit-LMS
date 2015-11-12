package mcpe

import "rockit/util/binary"

//RemoveEntityPacket is a packet implements <TODO>
type RemoveEntityPacket struct{}

//Encode encodes the packet
func (pk *RemoveEntityPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk RemoveEntityPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
