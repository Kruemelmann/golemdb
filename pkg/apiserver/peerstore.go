package apiserver

import (
	"errors"
	"sync"
)

type Peer struct {
	Id    string
	Route string
}

type PeerStore struct {
	registeredPeers []Peer
	mutex           sync.Mutex
}

func newPeerStore() *PeerStore {
	return &PeerStore{
		registeredPeers: []Peer{},
		mutex:           sync.Mutex{},
	}
}

func (p *PeerStore) Add(id, route string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// check if already in peerstore
	_, err := p.findIndexById(id)
	if err == nil {
		return errors.New("Peer with id: [" + id + "] is already in store")
	}
	_, err = p.findIndexByRoute(route)
	if err == nil {
		return errors.New("Peer with route: " + route + " is already in store")
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

	index, err := p.findIndexById(id)
	if err != nil {
		return errors.New("Peer with id: [" + id + "] could not be found in store")
	}

	p.registeredPeers = append(p.registeredPeers[:index], p.registeredPeers[index+1:]...)
	return nil
}

func (p *PeerStore) GetAll() ([]Peer, error) {
	return p.registeredPeers, nil
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
