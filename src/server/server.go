//Package server contains main server code
package server

import (
	"network/socket"
	"sync"
)

//Server is a struct with server-specific informations
type Server struct {
	ServerID uint64
	Socket   socket.Socket
}

func suspend() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}

//Start initializes server and starts
func (server *Server) Start() {
	socket.ServerID = server.ServerID
	go server.Socket.ProcessSend()
	go server.Socket.ProcessRecv()
	suspend()
}
