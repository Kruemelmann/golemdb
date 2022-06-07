package main

import (
	"fmt"

	"github.com/kruemelmann/golemdb/pkg/peerserver"
	"github.com/kruemelmann/golemdb/pkg/raft"
)

func main() {
	fmt.Println("Golem starting")

	//grpc_server, kv_server :=
	raft.NewConsensusModule()

	//FIXME on this point init the grpc server
	srv := peerserver.NewApiServer()
	srv.Start()
}
