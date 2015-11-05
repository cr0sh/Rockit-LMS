package protocol

import (
	"fmt"
	"network/packet"
)

var packetPool map[byte]packet.MCPEPacket

//RegisterPacket adds MCPE Packet to pool with given header
func RegisterPacket(head byte, pk packet.MCPEPacket) {
	if _, ok = packetPool[head]; ok {
		return
	}
	packetPool[head] = pk
}

//GetPacket gets MCPE Packet from pool with given header, and returns it
func GetPacket(head byte) (pk packet.MCPEPacket, err error) {
	if pk, ok := packetPool[head]; !ok {
		return pk, fmt.Errorf("Unimplemented or unregistered packet")
	}
	return pk, nil
}
