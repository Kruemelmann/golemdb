package raft

func NewConsensusModule() *ConsensusModule {
	return &ConsensusModule{
		State: State.Follower,
	}
}

type ConsensusModule struct {
	State StateType
}

func (c *ConsensusModule) RequestVotes()  {}
func (c *ConsensusModule) AppendEntries() {}
