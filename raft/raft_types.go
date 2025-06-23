package raft

import (
	"sync"
	"time"
)

type NodeState int

const (
	Follower  NodeState = iota // Passive node, waits for the instruction like(logs) form the leader (heartbeat)
	Candidate                  // Tries to get elected as Leader if it hasn’t heard from a Leader in time
	Leader                     // The boss node, sends instruction (logs) to the followers (heartbeat)
)

// iota makes these states integer constants (0, 1, 2)

type LogEntry struct {
	Term    int         // When this command was added
	Command interface{} // Command like key-value pair
}

type RaftNode struct {
	sync.Mutex
	Id          string
	State       NodeState
	CurrentTerm int
	VotedFor    string // Who this node voted for in the current term
	Logs        []LogEntry

	// Volatile state
	CommitIndex int // Index of the highest log entry known to be committed
	LastApplied int // Index of the highest log entry applied to state machine

	// Leader state
	NextIndex  map[string]int // For each follower, index of the next log entry to send to that follower
	MatchIndex map[string]int // For each follower, the highest index known to be replicated

	// Cluster state
	Peers map[string]string // Map of tcp addresses

	// Channels for communication
	RequestVoteChan chan bool // Channel for receiving request votes
	AppendEntriesChan chan bool // Channel for receiving append entries or heartbeat

	// Election timeout
	ElectionTimer time.Duration //  If no heartbeat is received within this, start an election
	HeartbeatTimer time.Duration // Interval between sending heartbeats (Leader to Followers)
}
