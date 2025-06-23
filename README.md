# Distributed Key-Value Database with Raft Consensus

A fault-tolerant distributed key-value storage system implementing the Raft consensus algorithm for ensuring data consistency across multiple nodes.

## 🚀 Features

- **Raft Consensus Algorithm**: Complete implementation with leader election, log replication, and fault tolerance
- **Multi-Node Cluster**: Support for distributed deployment across multiple servers
- **Fault Tolerance**: Handles node failures and network partitions gracefully
- **User Authentication**: Secure user registration and session management
- **Concurrent Operations**: Support for multiple simultaneous client connections
- **CRUD Operations**: Complete key-value store functionality (SET, GET, HAS, DELETE)
- **TCP Networking**: Custom protocol over TCP for client-server communication

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Raft Node 1   │    │   Raft Node 2   │    │   Raft Node 3   │
│   (Leader)      │◄──►│   (Follower)    │◄──►│   (Follower)    │
│                 │    │                 │    │                 │
│ ┌─────────────┐ │    │ ┌─────────────┐ │    │ ┌─────────────┐ │
│ │ Key-Value   │ │    │ │ Key-Value   │ │    │ │ Key-Value   │ │
│ │ Store       │ │    │ │ Store       │ │    │ │ Store       │ │
│ └─────────────┘ │    │ └─────────────┘ │    │ └─────────────┘ │
└─────────────────┘    └─────────────────┘    └─────────────────┘
        ▲                        ▲                        ▲
        │                        │                        │
        └────────── Clients ──────────────────────────────┘
```

## 🛠️ Technology Stack

- **Language**: Go (Golang)
- **Consensus Algorithm**: Raft
- **Networking**: TCP/IP, RPC
- **Concurrency**: Goroutines, Channels, Mutexes
- **Testing**: Go testing, Python automation
- **Architecture**: Distributed Systems, Microservices

## 📁 Project Structure

```
├── main.go              # Entry point for single node server
├── cache/               # Cache interface and implementation
│   ├── cache_interface.go
│   └── cache_impl.go
├── raft/               # Raft consensus implementation
│   ├── raft_types.go   # Data structures and types
│   ├── raft_node.go    # Core Raft node logic
│   └── raft_rpc.go     # RPC handlers for consensus
├── server/             # Server implementation
│   └── server.go       # TCP server and client handling
├── store/              # User data storage
│   └── store.go        # User store implementation
└── test_client.py      # Testing and validation scripts
```

## 🚦 Getting Started

### Prerequisites

- Go 1.21 or higher
- Python 3.x (for testing)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Mr-ut/Distributed-Key-Value-DB.git
cd Distributed-Key-Value-DB
```

2. Install dependencies:
```bash
go mod tidy
```

### Running a Single Node

```bash
go run main.go
```

The server will start on port 6000. Connect using:
```bash
telnet localhost 6000
```

### Testing the System

Run the basic functionality tests:
```bash
python3 test_client.py
```

## 📖 Usage

### Client Commands

After connecting to a server:

1. **Authentication**:
   ```
   SIGNUP
   <username>
   <password>
   ```

2. **Key-Value Operations**:
   ```
   SET key value     # Store a key-value pair
   GET key          # Retrieve value for key
   HAS key          # Check if key exists
   DELETE key       # Remove key-value pair
   ```

### Example Session

```bash
$ telnet localhost 6000
Connected to localhost.
Escape character is '^]'.
LOGIN or SIGNUP
SIGNUP
Enter username: alice
Enter password: secret123
User created successfully
SET mykey myvalue
OK
GET mykey
myvalue
HAS mykey
1
DELETE mykey
OK
```

## 🧪 Testing

### Basic Functionality Test
```bash
python3 test_client.py
```

### Manual Testing
You can test multiple clients simultaneously by opening multiple telnet sessions:
```bash
# Terminal 1
telnet localhost 6000

# Terminal 2  
telnet localhost 6000
```

## 🔧 Raft Consensus Implementation

### Key Components

1. **Leader Election**: Randomized election timeouts prevent split votes
2. **Log Replication**: Leaders replicate entries to followers before committing
3. **Safety**: Only servers with up-to-date logs can become leaders
4. **Fault Tolerance**: Majority consensus ensures availability during failures

### Raft States

- **Follower**: Passive state, responds to leader and candidate requests
- **Candidate**: Tries to become leader during election
- **Leader**: Handles client requests and replicates log entries

## 🏗️ System Design Principles

### CAP Theorem Compliance
- **Consistency**: Strong consistency through Raft consensus
- **Availability**: Available when majority of nodes are operational
- **Partition Tolerance**: Handles network partitions gracefully

### Fault Tolerance
- Tolerates up to `(n-1)/2` node failures in an `n`-node cluster
- Automatic leader re-election on failure
- Network partition recovery

## 🚀 Performance Characteristics

- **Consistency**: 100% strong consistency via Raft protocol
- **Availability**: High availability with majority node operation
- **Latency**: Low latency for read operations from leader
- **Throughput**: Scales with cluster size for read operations

## 🔮 Future Enhancements

- [ ] Persistent storage with disk-based logs
- [ ] Snapshot mechanism for log compaction
- [ ] HTTP REST API interface
- [ ] Metrics and monitoring dashboard
- [ ] Configuration management
- [ ] Read replicas for improved read performance
- [ ] Encryption for client-server communication

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Raft Consensus Algorithm](https://raft.github.io/) by Diego Ongaro and John Ousterhout
- [The Raft Paper](https://raft.github.io/raft.pdf) for algorithm specification
- Go community for excellent concurrency primitives

## 📧 Contact

- GitHub: [@Mr-ut](https://github.com/Mr-ut)
- Project Link: [https://github.com/Mr-ut/Distributed-Key-Value-DB](https://github.com/Mr-ut/Distributed-Key-Value-DB)

---

⭐ **Star this repository if you found it helpful!**