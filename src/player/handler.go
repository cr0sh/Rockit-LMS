package player

import (
	"net"
	"network/packet/mcpe"
)

//Handler handles packets from player and controls player entity.
type Handler struct {
	Address  net.UDPAddr
	username string
}

//HandlePacket handles MCPE DataPacket from player
func (handler *Handler) HandlePacket(pk []byte) {
	mcpe.GetPacket(pk[0])
}
