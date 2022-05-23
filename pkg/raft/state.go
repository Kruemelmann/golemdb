package raft

type StateType int

var State = StateFactory()

func StateFactory() *states {
	return &states{
		Follower:  1,
		Candidate: 2,
		Leader:    3,
	}
}

type states struct {
	Follower  StateType
	Candidate StateType
	Leader    StateType
}
