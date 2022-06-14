package peerserver

import (
	"log"
	"net"
	"sync"

	pb "github.com/kruemelmann/golemdb/pkg/api/pb/peers"
	"github.com/kruemelmann/golemdb/pkg/peerserver/peerstore"
	"google.golang.org/grpc"
)

var peerserverInstance *PeerServer
var once sync.Once

type PeerServer struct {
	pb.UnimplementedPeersServiceServer
	PeerStore *peerstore.PeerStore
}

func NewPeerServer() *PeerServer {
	once.Do(func() {
		peerserverInstance = &PeerServer{
			PeerStore: peerstore.NewPeerStore(),
		}
		//initialy fill the peerstore
		peerserverInstance.InitialRegisterPeers()
	})
	return peerserverInstance
}

func (a *PeerServer) InitialRegisterPeers() {
	//FIXME remove debugging peers and read them from env or config file
	a.PeerStore.Add("123", "127.0.0.1:9091")
}

func (a *PeerServer) Start() {
	//apiaddress := viper.GetString("registry.host") + ":" + viper.GetString("registry.port")
	//binding ports
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Starting grpc service for peers\n")
	grpcserver := grpc.NewServer()
	pb.RegisterPeersServiceServer(grpcserver, &PeerServer{})
	grpcserver.Serve(lis)
}

func (a *PeerServer) ListPeerIds() ([]string, error) {
	ids := []string{}

	p, err := a.PeerStore.GetAll()
	if err != nil {
		log.Fatalf("Peerstore getall error %s", err.Error())
		return nil, err
	}
	for _, v := range p {
		ids = append(ids, v.Id)
	}

	return ids, nil
}

type RequestVoteArgs struct {
	Term        int
	CandidateId string
}
type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

//TODO
func (a *PeerServer) RequestVote(id string, args RequestVoteArgs) (*RequestVoteReply, error) {
	return nil, nil
}
