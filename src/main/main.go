package main

import (
	"fmt"
	"math/rand"
	"network/socket"
	"os"
	"server"
)

func main() {
	server := server.Server{uint64(rand.Uint32()), *new(socket.Socket)}
	if err := server.Socket.Open(19132); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server.Start()
}
