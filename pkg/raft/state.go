package raft

type StateType int

var State = StateFactory()

func StateFactory() *states {
	return &states{
		Follower:  1,
		Candidate: 2,
		Leader:    3,
		Dead:      4,
	}
}

type states struct {
	Follower  StateType
	Candidate StateType
	Leader    StateType
	Dead      StateType
}

func (s StateType) String() string {
	switch s {
	case 1:
		return "Follower"
	case 2:
		return "Candidate"
	case 3:
		return "Leader"
	case 4:
		return "Dead"
	default:
		return "Invalid State"
	}

}
