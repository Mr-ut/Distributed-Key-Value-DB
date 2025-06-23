package raft

type RequestVoteArgs struct {
	Term         int    // Candidate’s term
	CandidateId  string // Candidate requesting vote
	LastLogIndex int    // Index of candidate's last log entry
	LastLogTerm  int    //
}

type RequestVoteResponse struct {
	Term        int  // Current term of node responding (useful for candidates to update their term)
	VoteGranted bool // True if candidate received vote
}

type AppendEntriesArgs struct {
	Term         int    // Leader’s term
	LeaderId     string // id of leader
	PrevLogIndex int    // for log consistency check
	PrevLogTerm  int
	Entries      []LogEntry // New entries for followers
	LeaderCommit int        // Leader's commitIndex
}

type AppendEntriesResponse struct {
	Term    int // current term of follower
	Success bool
}

func (n *RaftNode) RequestVote(request *RequestVoteArgs, response *RequestVoteResponse) error {
	n.Lock() // Lock the node to prevent race conditions
	defer n.Unlock()

    if request.Term < n.CurrentTerm { // If req term is older than current term, reject vote
		response.Term = n.CurrentTerm
		response.VoteGranted = false
		return nil
	}

	if request.Term > n.CurrentTerm {
		n.CurrentTerm = request.Term
		n.State = Follower // If req term is newer, update current term and become follower
		n.VotedFor = ""
	}

	lastLogIndex := len(n.Logs) - 1
	lastLogTerm := 0

	if lastLogIndex >= 0 {
		lastLogTerm = n.Logs[lastLogIndex].Term // Get the term of the last log entry
	}

    upToDate := request.LastLogIndex > lastLogTerm || (request.LastLogTerm == n.CurrentTerm && request.LastLogIndex >= lastLogIndex)

	if (n.VotedFor == "" || n.VotedFor == request.CandidateId) && upToDate {
		response.VoteGranted = true
		n.VotedFor = request.CandidateId
		n.ResetElectionTimer()
	}else{
		response.VoteGranted = false
	}
	response.Term = n.CurrentTerm
	return nil
}


func (n *RaftNode) AppendEntries(request *AppendEntriesArgs, response *AppendEntriesResponse) error {
	n.Lock()
	defer n.Unlock()

	response.Success = false
	response.Term = n.CurrentTerm

	if request.Term < n.CurrentTerm { // If request term is older, reject
		return nil
	}

	if request.Term > n.CurrentTerm {
		n.CurrentTerm = request.Term
		n.State = Follower
		n.VotedFor = ""
	}

	 n.ResetElectionTimer()
	if request.PrevLogIndex >= 0 && n.Logs[request.PrevLogIndex].Term != request.PrevLogTerm {
		return nil
	}

	n.Logs = n.Logs[:request.PrevLogIndex+1] // Truncate logs to the previous index
	n.Logs = append(n.Logs, request.Entries...) // Append new entries

	if request.LeaderCommit > n.CommitIndex {
		n.CommitIndex = min(request.LeaderCommit, len(n.Logs) - 1)
	}

	response.Success = true
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
