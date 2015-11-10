package mcpe

import "bytes"

//AdventureSettingsPacket is a packet implements <TODO>
type AdventureSettingsPacket struct{}

//Encode encodes the packet
func (pk *AdventureSettingsPacket) Encode(fields map[string]interface{}) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AdventureSettingsPacket) Decode(buf *bytes.Buffer) (fields map[string]interface{}, err error) {
	return
}
