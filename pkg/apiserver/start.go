package apiserver

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type ApiServer struct {
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

	grpcserver := grpc.NewServer()
	NewPeerServer(grpcserver)
	grpcserver.Serve(lis)
}
