package peerserver

import (
	"log"
	"net"

	pb "github.com/kruemelmann/golemdb/pkg/api/pb/peers"
	"github.com/kruemelmann/golemdb/pkg/peerserver/peerstore"
	"google.golang.org/grpc"
)

type ApiServer struct {
	pb.UnimplementedPeersServiceServer
	PeerStore peerstore.PeerStore
}

func NewApiServer() *ApiServer {
	//TODO make it singelton
	return &ApiServer{}
}

func (a *ApiServer) Start() {
	//apiaddress := viper.GetString("registry.host") + ":" + viper.GetString("registry.port")
	//binding ports
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Starting grpc service for peers\n")
	grpcserver := grpc.NewServer()
	pb.RegisterPeersServiceServer(grpcserver, &ApiServer{})
	grpcserver.Serve(lis)
}
