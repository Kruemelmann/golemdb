package raft

type RequestVoteArgs struct {
	Term        int
	CandidateId string
}
type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

//TODO this is only debugging
func RequestVotes(id string, args RequestVoteArgs) (RequestVoteReply, error) {
	reply := RequestVoteReply{
		Term:        args.Term,
		VoteGranted: true,
	}
	return reply, nil
}

type AppendEntriesArgs struct {
	Term     int
	LeaderId string
}
type AppendEntriesReply struct {
	Term int
}

func AppendEntries(id string, args AppendEntriesArgs) (AppendEntriesReply, error) {
	reply := AppendEntriesReply{
		Term: args.Term,
	}
	//TODO
	return reply, nil
}
