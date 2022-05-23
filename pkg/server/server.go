package server

import "sync"

var (
	instance *GolemServer
	once     sync.Once
)

type GolemServer struct {
}

func NewGolemServer() *GolemServer {
	once.Do(func() {
		instance = &GolemServer{}
	})
	return instance
}
