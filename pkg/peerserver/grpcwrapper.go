package peerserver

import pb "github.com/kruemelmann/golemdb/pkg/api/pb/peers"

type grpcWrapper struct {
	pb.UnimplementedPeersServiceServer
}
