package apiserver

import (
	"errors"
	"log"
	"sync"
)

var peerstoreInstance *PeerStore
var peerstoreOnce sync.Once

func newPeerStore() *PeerStore {
	once.Do(func() {
		peerstoreInstance = &PeerStore{
			registeredPeers: []Peer{},
			mutex:           sync.Mutex{},
		}
	})
	return peerstoreInstance
}

type Peer struct {
	Id    string
	Route string
}
type PeerStore struct {
	registeredPeers []Peer
	mutex           sync.Mutex
}

func (p *PeerStore) Add(id, route string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// check if already in peerstore
	_, err := p.findIndexById(id)
	if err == nil {
		log.Fatalf("Peer with id: [%s] is already in store", id)
	}
	_, err = p.findIndexByRoute(route)
	if err == nil {
		log.Fatalf("Peer with route: %s is already in store", route)
	}

	// append to store
	p.registeredPeers = append(p.registeredPeers, Peer{
		Id:    id,
		Route: route,
	})

	return nil
}

func (p *PeerStore) Remove(id string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return nil
}

func (p *PeerStore) GetById(id string) (*Peer, error) {
	index, err := p.findIndexById(id)
	if err != nil {
		return nil, err
	}
	return &p.registeredPeers[index], nil
}

func (p *PeerStore) GetByRoute(route string) (*Peer, error) {
	index, err := p.findIndexByRoute(route)
	if err != nil {
		return nil, err
	}
	return &p.registeredPeers[index], nil
}

//======== PeerStore Utilfunctions
func (p *PeerStore) findIndexById(id string) (int, error) {
	for k, v := range p.registeredPeers {
		if v.Id == id {
			return k, nil
		}
	}
	return -1, errors.New("Peer with id: [" + id + "] could not be found in store")
}
func (p *PeerStore) findIndexByRoute(route string) (int, error) {
	for k, v := range p.registeredPeers {
		if v.Route == route {
			return k, nil
		}
	}
	return -1, errors.New("Peer with route:" + route + " could not be found in store")
}
