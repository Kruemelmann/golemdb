package raft

type ConsensusModule struct {
	State StateType
}

func NewConsensusModule() *ConsensusModule {
	return &ConsensusModule{
		State: State.Follower,
	}
}

func (c *ConsensusModule) RequestVotes()  {}
func (c *ConsensusModule) AppendEntries() {}
