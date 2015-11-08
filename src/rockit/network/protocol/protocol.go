//Package protocol defines protocol constants, like versions.
package protocol

const (
	//RaknetProtocol defines current RakNet version
	RaknetProtocol = 7
	//RaknetMagic defines special magic value used for Raknet
	RaknetMagic = "\x00\xff\xff\x00\xfe\xfe\xfe\xfe\xfd\xfd\xfd\xfd\x12\x34\x56\x78"
)
