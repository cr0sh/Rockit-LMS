package mcpe

import "rockit/util/binary"

//LoginPacket is a packet implements MCPE login packet
type LoginPacket struct{}

//Encode encodes the packet
func (pk *LoginPacket) Encode(fields Field) (buf []byte, err error) {
	return
}

//Decode decodes the packet
func (pk LoginPacket) Decode(buf binary.Stream) (fields Field, err error) {
	fields = make(Field)
	fields["username"] = buf.ReadString()
	fields["protocol1"] = buf.ReadInt()
	fields["protocol2"] = buf.ReadInt()
	if fields["protocol1"].(uint32) < Protocol {
		return
	}
	fields["clientID"] = buf.ReadLong()
	fields["UUID"] = buf.Read(16)
	fields["serverAddr"] = buf.ReadString()
	fields["clientSecret"] = []byte(buf.ReadString())
	fields["slimness"] = buf.Read(1)[0] > 0
	fields["skin"] = []byte(buf.ReadString())
	return
}
