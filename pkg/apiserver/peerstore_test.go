package apiserver_test

import (
	"sync"
	"testing"

	"github.com/kruemelmann/golemdb/pkg/apiserver"
)

var testinstance *apiserver.PeerStore
var once sync.Once

func init() {
	once.Do(func() {
		testinstance = &apiserver.PeerStore{}
	})
	testinstance.Add("123", "127.0.0.1:9090")
}

func TestNewPeerStore(t *testing.T) {
	t.Run("add a peer to the store", func(t *testing.T) {
		testinstance.Add("234", "127.0.0.1:9091")
	})
	t.Run("get a peer from the store", func(t *testing.T) {
		p, err := testinstance.GetById("123")
		if err != nil {
			t.Errorf("Failed with error %s", err.Error())
		}

		if p.Id != "123" {
			t.Errorf("Failed with wrong id")
		}
	})
	t.Run("remove a peer from the store", func(t *testing.T) {
		testinstance.Add("345", "127.0.0.1:9092")

		err := testinstance.Remove("345")
		if err != nil {
			t.Errorf("Error while removing %s", err)
		}
		err = testinstance.Remove("345")
		if err == nil {
			t.Errorf("Error removing should not work twice err: %s", err)
		}
	})
}
