package mcpe

import "rockit/util/binary"

//DisconnectPacket is a packet implements <TODO>
type DisconnectPacket struct{}

//Encode encodes the packet
func (pk *DisconnectPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk DisconnectPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
