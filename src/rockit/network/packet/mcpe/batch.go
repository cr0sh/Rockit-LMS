package mcpe

import (
	"fmt"
	"rockit/util/binary"
)

//BatchPacket is a packet implements batch packet
type BatchPacket struct{}

//Encode encodes the packet
func (pk *BatchPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk BatchPacket) Decode(buf binary.Stream) (fields Field, err error) {
	fields = make(Field)
	size := buf.ReadInt()
	if size == 0 {
		err = fmt.Errorf("Invalid payload size: 0")
		return
	}
	fields["payload"] = buf.Read(uint(size))
	fields["size"] = size
	return
}
