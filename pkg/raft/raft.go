package raft

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

func NewConsensusModule() *ConsensusModule {
	c := ConsensusModule{}
	c.localTerm = 1
	c.state = State.Follower
	c.ID = uuid.New().String()

	go func() {
		c.mutex.Lock()
		c.lastElectionReset = time.Now()
		c.mutex.Unlock()
		c.startElectionTimer()
	}()

	return &c
}

type ConsensusModule struct {
	ID    string
	mutex sync.Mutex
	state StateType

	localTerm int
	//election
	lastElectionReset time.Time
	votedId           string
}

func (c *ConsensusModule) startElectionTimer() {
	rand.Seed(time.Now().UnixNano())
	timeout := time.Duration(400+rand.Intn(100)) * time.Millisecond
	startterm := c.localTerm
	log.Printf("election timer started at %v term: %d", timeout, startterm)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		<-ticker.C

		c.mutex.Lock()
		if c.state != State.Follower && c.state != State.Candidate {
			log.Printf("stop election timer on state %s\n", c.state.String())
			c.mutex.Unlock()
			return
		}

		if d := time.Since(c.lastElectionReset); d >= timeout {
			c.startElection()
			c.mutex.Unlock()
			return
		}
		c.mutex.Unlock()
	}
}

func (c *ConsensusModule) startElection() {
	c.state = State.Candidate
	c.lastElectionReset = time.Now()
	c.votedId = c.ID
	log.Printf("[%s] started election on state %s\n", c.ID, c.state.String())
}
