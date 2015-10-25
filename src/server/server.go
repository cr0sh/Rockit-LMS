package server

import (
	"network/socket"
	"sync"
)

type Server struct {
	ServerID uint64
	Socket   socket.Socket
}

func suspend() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Wait()
}

func (server *Server) Start() {
	socket.ServerID = server.ServerID
	go server.Socket.ProcessSend()
	go server.Socket.ProcessRecv()
	suspend()
}
