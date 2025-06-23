package raft

import (
	"math/rand"
	"net/rpc"
	"sync"
	"time"
)

func NewRaftNode(id string, peers map[string]string) *RaftNode {
	node := &RaftNode{
		Id:                id,
		State:             Follower,
		CurrentTerm:       0,
		VotedFor:          "",
		Logs:              []LogEntry{{Term: 0, Command: nil}}, // Genesis Entry
		CommitIndex:       -1,
		LastApplied:       -1,
		Peers:             peers,
		NextIndex:         make(map[string]int),
		MatchIndex:        make(map[string]int),
		RequestVoteChan:   make(chan bool),
		AppendEntriesChan: make(chan bool),
	}

	return node
}

func (n *RaftNode) ResetElectionTimer() {
	// Note: This is a placeholder. In a real implementation, you'd want to use an actual timer
	// For now, we'll just set the duration
	timeout := time.Duration(150+rand.Intn(150)) * time.Millisecond
	n.ElectionTimer = timeout
	// In a real implementation, you'd start a timer here that calls startElection after timeout
}

func (n *RaftNode) startElection() {
	n.Lock()
	if n.State == Leader {
		n.Unlock()
		return
	}
	n.State = Candidate
	n.CurrentTerm++
	n.VotedFor = n.Id
	currentTerm := n.CurrentTerm

	// Prepare RequestVote rpc
	args := &RequestVoteArgs{
		Term:         currentTerm,
		CandidateId:  n.Id,
		LastLogIndex: len(n.Logs) - 1,
		LastLogTerm:  -1,
	}
	if args.LastLogIndex >= 0 {
		args.LastLogTerm = n.Logs[args.LastLogIndex].Term
	}
	n.Unlock()

	// Send RequestVote RPCs to all peers
	votes := 1 // Vote for self
	var voteMutex sync.Mutex

	for peerAddr := range n.Peers {
		if peerAddr == n.Id {
			continue
		}
		go func(peerAddr string) {
			var reply RequestVoteResponse
			client, err := rpc.Dial("tcp", peerAddr)
			if err != nil {
				return
			}
			defer client.Close()
			if err := client.Call("RaftNode.RequestVote", args, &reply); err != nil {
				return
			}
			n.Lock()
			if reply.Term > n.CurrentTerm {
				n.CurrentTerm = reply.Term
				n.State = Follower
				n.VotedFor = ""
				n.Unlock()
				return
			}
			n.Unlock()
			if reply.VoteGranted {
				voteMutex.Lock()
				votes++
				if votes > len(n.Peers)/2 && n.State == Candidate {
					n.Lock()
					if n.State == Candidate {
						n.State = Leader
						for peerAddr := range n.Peers {
							n.NextIndex[peerAddr] = len(n.Logs)
							n.MatchIndex[peerAddr] = -1
						}
						n.StartHeartbeat()
					}
					n.Unlock()
				}
				voteMutex.Unlock()
			}
		}(peerAddr)
	}
}

func (n *RaftNode) StartHeartbeat() {
	// Note: This is a placeholder. In a real implementation, you'd want to use an actual timer
	// For now, we'll just set the duration and call sendHeartbeat
	n.HeartbeatTimer = 100 * time.Millisecond
	n.sendHeartbeat()
	// In a real implementation, you'd start a timer here that calls sendHeartbeat periodically
}

func (n *RaftNode) sendHeartbeat() {
	n.Lock()
	if n.State != Leader {
		n.Unlock()
		return
	}
	for peerAddr := range n.Peers {
		if peerAddr == n.Id {
			continue
		}
		go func(peerAddr string) {
			for {
				n.Lock()
				prevLogIndex := n.NextIndex[peerAddr] - 1
				prevLogTerm := 0
				if prevLogIndex >= 0 {
					prevLogTerm = n.Logs[prevLogIndex].Term
				}

				entries := n.Logs[n.NextIndex[peerAddr]:]
				args := AppendEntriesArgs{
					Term:         n.CurrentTerm,
					LeaderId:     n.Id,
					PrevLogIndex: prevLogIndex,
					PrevLogTerm:  prevLogTerm,
					Entries:      entries,
					LeaderCommit: n.CommitIndex,
				}
				n.Unlock()
				var reply AppendEntriesResponse
				client, err := rpc.DialHTTP("tcp", peerAddr)
				if err != nil {
					return
				}
				defer client.Close()
				if err := client.Call("RaftNode.AppendEntries", args, &reply); err != nil {
					return
				}
				n.Lock()
				if reply.Term > n.CurrentTerm {
					n.CurrentTerm = reply.Term
					n.State = Follower
					n.VotedFor = ""
					n.Unlock()
					return
				}
				if reply.Success {
					n.NextIndex[peerAddr] = prevLogIndex + len(entries) + 1
					n.MatchIndex[peerAddr] = n.NextIndex[peerAddr] - 1
					n.Unlock()
					break
				} else {
					if n.NextIndex[peerAddr] > 0 {
						n.NextIndex[peerAddr]--
					}
					n.Unlock()
				}
			}
		}(peerAddr)
	}
	n.Unlock()
}
