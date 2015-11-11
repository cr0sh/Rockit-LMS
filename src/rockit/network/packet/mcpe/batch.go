package mcpe

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//BatchPacket is a packet implements <TODO>
type BatchPacket struct{}

//Encode encodes the packet
func (pk *BatchPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk BatchPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	fields = make(Field)
	size := new(uint32)
	binary.Read(buf, binary.BigEndian, size)
	if *size == 0 {
		err = fmt.Errorf("Invalid payload size: 0")
		return
	}
	fields["payload"] = buf.Next(int(*size))
	fields["size"] = *size
	return
}
