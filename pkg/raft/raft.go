package raft

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kruemelmann/golemdb/pkg/peerserver"
)

func NewConsensusModule() *ConsensusModule {
	c := ConsensusModule{}
	c.currentTerm = 1
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

	currentTerm int
	//election
	lastElectionReset time.Time
	votedId           string
}

func (c *ConsensusModule) startElectionTimer() {
	rand.Seed(time.Now().UnixNano())
	timeout := time.Duration(400+rand.Intn(100)) * time.Millisecond
	c.mutex.Lock()
	startterm := c.currentTerm
	c.mutex.Unlock()
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

		if startterm != c.currentTerm {
			log.Printf("stop election term changed from %d to %d\n", startterm, c.currentTerm)
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
	c.currentTerm++
	savedCurrentTerm := c.currentTerm
	c.lastElectionReset = time.Now()
	c.votedId = c.ID
	log.Printf("[%s] started election on state %s with term %d\n", c.ID, c.state.String(), savedCurrentTerm)
	votesReceived := 1

	//TODO handle error
	ids, _ := peerserver.NewPeerServer().ListPeerIds()
	for _, id := range ids {
		go func(ID string) {
			args := RequestVoteArgs{
				Term:        savedCurrentTerm,
				CandidateId: c.ID,
			}

			reply, err := RequestVotes(ID, args)
			if err == nil {
				c.mutex.Lock()
				defer c.mutex.Unlock()
				log.Printf("received RequestVoteReply term=%v bool=%v\n", reply.Term, reply.VoteGranted)

				if c.state != State.Candidate {
					log.Printf("while waiting for reply, state = %v\n", c.state.String())
					return
				}

				if reply.Term > savedCurrentTerm {
					log.Printf("term out of date in RequestVoteReply\n")
					c.becomeFollower(reply.Term)
					return
				} else if reply.Term == savedCurrentTerm {
					if reply.VoteGranted {
						votesReceived++
						if votesReceived*2 > len(ids)+1 {
							// Won the election!
							log.Printf("wins election with %d votes\n", votesReceived)
							c.startLeader()
							return
						}
					}
				}
			}
		}(id)
	}
	go c.startElectionTimer()
}

func (c *ConsensusModule) startLeader() {
	c.state = State.Leader
	log.Printf("becomes Leader; term=%d", c.currentTerm)
	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		for {
			c.leaderSendHeartbeats()
			<-ticker.C

			c.mutex.Lock()
			if c.state != State.Leader {
				c.mutex.Unlock()
				return
			}
			c.mutex.Unlock()
		}
	}()
}

func (c *ConsensusModule) leaderSendHeartbeats() {
	c.mutex.Lock()
	savedCurrentTerm := c.currentTerm
	log.Printf("-> Term %d State %s", savedCurrentTerm, c.state.String())
	c.mutex.Unlock()

	ids, _ := peerserver.NewPeerServer().ListPeerIds()
	for _, id := range ids {
		args := AppendEntriesArgs{
			Term:     savedCurrentTerm,
			LeaderId: c.ID,
		}
		go func(id string) {
			reply, err := AppendEntries(id, args)
			if err != nil {
				log.Printf("Error while AppendEntries Request on id %s\n", id)
			}
			c.mutex.Lock()
			defer c.mutex.Unlock()
			if reply.Term > savedCurrentTerm {
				log.Printf("term out of date in heartbeat reply\n")
				c.becomeFollower(reply.Term)
				return
			}
		}(id)
	}
}

func (c *ConsensusModule) becomeFollower(term int) {
	log.Printf("becomes Leader; term=%d", c.currentTerm)
	c.state = State.Follower
	c.currentTerm = term
	c.votedId = ""
	c.lastElectionReset = time.Now()

	go c.startElectionTimer()
}
