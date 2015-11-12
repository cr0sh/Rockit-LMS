package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"rockit/network"
	"rockit/server"
	"rockit/util/logger"
	"runtime"
	"runtime/trace"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	dbg := flag.Bool("d", false, "sets loglevel to 0(debug) if set")
	tr := flag.Bool("t", false, "prints execution trace log to Stdout - DO NOT USE NOW")
	flag.Parse()
	if *dbg {
		logger.SetLevel(0)
	}
	if *tr {
		trace.Start(os.Stdout)
	}
	server := server.Server{ServerID: uint64(rand.Uint32()), Socket: *new(network.Socket)}
	if err := server.Socket.Open(19132); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server.Start()
}
