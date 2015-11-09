package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"rockit/network"
	"rockit/server"
	"rockit/util"
)

func main() {
	debug := flag.Bool("d", false, "sets loglevel to 0(debug) if set")
	flag.Parse()
	if *debug {
		util.SetLevel(0)
	}
	server := server.Server{ServerID: uint64(rand.Uint32()), Socket: *new(network.Socket)}
	if err := server.Socket.Open(19132); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server.Start()
}
