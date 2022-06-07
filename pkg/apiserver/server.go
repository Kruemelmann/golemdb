package apiserver

import (
	"sync"

	pb "github.com/kruemelmann/golemdb/pkg/api/pb/peers"
	"google.golang.org/grpc"
)

var instance *PeerServer
var once sync.Once

func NewPeerServer(grpcserver *grpc.Server) {
	pb.RegisterPeersServiceServer(grpcserver, newServer())
}

func newServer() *PeerServer {
	once.Do(func() {
		instance = &PeerServer{}
	})
	return instance
}

type PeerServer struct {
	pb.UnimplementedPeersServiceServer
	PeerStore PeerStore
}
