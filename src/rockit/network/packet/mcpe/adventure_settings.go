package mcpe

import "rockit/util/binary"

//AdventureSettingsPacket is a packet implements <TODO>
type AdventureSettingsPacket struct{}

//Encode encodes the packet
func (pk *AdventureSettingsPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk AdventureSettingsPacket) Decode(buf binary.Stream) (fields Field, err error) {
	return
}
