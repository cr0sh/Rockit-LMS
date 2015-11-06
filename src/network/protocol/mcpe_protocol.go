package protocol

import (
	"fmt"
	"player"
)

//MCPEPacket is a data packet interface, for MCPE Clients
type MCPEPacket interface {
	Encode() error
	Decode(*player.Player) error
	Fields() map[string]interface{}
}

var packetPool map[byte]MCPEPacket

//RegisterPacket adds MCPE Packet to pool with given header
func RegisterPacket(head byte, pk MCPEPacket) {
	if _, ok := packetPool[head]; ok {
		return
	}
	packetPool[head] = pk
}

//GetPacket gets MCPE Packet from pool with given header, and returns it
func GetPacket(head byte) (pk MCPEPacket, err error) {
	if pk, ok := packetPool[head]; !ok {
		return pk, fmt.Errorf("Unimplemented or unregistered packet")
	}
	return pk, nil
}
