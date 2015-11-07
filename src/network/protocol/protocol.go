package protocol

var (
	//RaknetProtocol defines current RakNet version
	RaknetProtocol = 7
	//RaknetMagic defines special magic value used for Raknet
	RaknetMagic = []byte("\x00\xff\xff\x00\xfe\xfe\xfe\xfe\xfd\xfd\xfd\xfd\x12\x34\x56\x78")
)
