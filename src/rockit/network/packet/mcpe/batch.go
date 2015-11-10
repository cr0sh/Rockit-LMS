package mcpe

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//BatchPacket is a packet implements <TODO>
type BatchPacket struct{}

//Encode encodes the packet
func (pk *BatchPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk BatchPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	fields = make(map[string]interface{})
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
