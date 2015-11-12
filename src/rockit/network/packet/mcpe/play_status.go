package mcpe

import "rockit/util/binary"

//PlayStatusPacket is a packet implements player status packet
type PlayStatusPacket struct{}

const (
	Success       = 0
	FailureClient = 1
	FailureServer = 2
	PlayerSpawn   = 3
)

//Encode encodes the packet
func (pk *PlayStatusPacket) Encode(fields Field) (buf []byte, err error) {
	stream := *new(binary.Stream)
	stream.WriteInt(fields["status"].(uint32))
	buf = stream.Buffer
	return
}

//Decode decodes the packet
func (pk PlayStatusPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
