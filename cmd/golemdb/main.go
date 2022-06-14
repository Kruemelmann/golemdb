package main

import (
	"fmt"
	"time"

	"github.com/kruemelmann/golemdb/pkg/peerserver"
	"github.com/kruemelmann/golemdb/pkg/raft"
)

func main() {
	fmt.Println("Golem starting")

	time.Sleep(10 * time.Second)
	//grpc_server, kv_server :=
	raft.NewConsensusModule()

	//FIXME on this point init the grpc server
	srv := peerserver.NewPeerServer()
	srv.Start()
}
