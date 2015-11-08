//Package server contains main server code
package server

import (
	"rockit/network"
	"rockit/player"
	"sync"
)

//Server is a struct with server-specific informations
type Server struct {
	ServerID   uint64
	Socket     network.Socket
	playerList map[uint]player.Player
}

func suspend() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}

//Start initializes server and starts
func (server *Server) Start() {
	network.ServerID = server.ServerID
	server.playerList = make(map[uint]player.Player)
	go server.Socket.ProcessSend()
	go server.Socket.ProcessRecv()
	suspend()
}
