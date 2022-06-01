package server

import (
	"sync"

	pb "github.com/kruemelmann/golemdb/pkg/api/pb/peers"
)

var instance *PeersService
var once sync.Once

func NewPeersServer() *PeersService {
	once.Do(func() {
		serverInstance = &PeersService{}
	})
	return instance
}

type PeersService struct {
	pb.UnimplementedPeersServiceServer
}
