package mcpe

import "bytes"

//AdventureSettingsPacket is a packet implements <TODO>
type AdventureSettingsPacket struct{}

//Encode encodes the packet
func (pk *AdventureSettingsPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AdventureSettingsPacket) Decode(buf *bytes.Buffer) (fields Field, err error) {
	return
}
