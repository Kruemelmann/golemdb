package main

import (
	"fmt"
	"time"

	"github.com/kruemelmann/golemdb/pkg/raft"
)

func main() {
	fmt.Println("Golem starting")

	//grpc_server, kv_server :=
	raft.NewConsensusModule()

	//FIXME on this point init the grpc server
	time.Sleep(10 * time.Second)
}
