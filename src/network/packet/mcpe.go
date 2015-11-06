package packet

import "player"

//MCPEPacket is a data packet interface, for MCPE Clients
type MCPEPacket interface {
	Encode() error
	Decode(*player.Player) error
	Fields() map[string]interface{}
}
