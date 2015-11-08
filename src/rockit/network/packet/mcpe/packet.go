//Package mcpe provides Mojang-defined MCPE packets
package mcpe

import (
	"fmt"
)

//Packet is a data packet interface, for MCPE Clients
type Packet interface {
	Encode() error
	Decode() error
	GetField(string) interface{}
	SetField(string) interface{}
}

var packetPool map[byte]Packet

//RegisterPacket adds MCPE Packet to pool with given header
func RegisterPacket(head byte, pk Packet) {
	if _, ok := packetPool[head]; ok {
		return
	}
	packetPool[head] = pk
}

//GetPacket gets MCPE Packet from pool with given header, and returns it
func GetPacket(head byte) (pk Packet, err error) {
	if pk, ok := packetPool[head]; !ok {
		return pk, fmt.Errorf("Unimplemented or unregistered packet")
	}
	return pk, nil
}
