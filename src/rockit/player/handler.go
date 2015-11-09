package player

import (
	"encoding/hex"
	"net"
	"rockit/network/packet/mcpe"
	"rockit/util"
)

//Handler handles packets from player and controls player entity.
type Handler struct {
	Address  net.UDPAddr
	entity   Player
	username string
}

//HandlePacket handles MCPE DataPacket from player
func (handler *Handler) HandlePacket(pk []byte) {
	util.Debug("Handling MCPE Packet: head 0x" + hex.EncodeToString([]byte{pk[0]}))
	var ppk mcpe.Packet
	var err error
	ppk, err = mcpe.GetPacket(pk[0])
	if err != nil {
		util.FromError(err, 1)
		return
	}
	ppk.Decode()
	switch pk[0] {

	}
}
