package main

import (
	"log"

	"github.com/Mr-ut/Distributed-Key-Value-DB/server"
)

func main() {
	// Create a new server on port 6000
	srv := server.NewServer(":6000")

	log.Println("Starting Distributed Key-Value Database Server...")
	log.Println("Server will listen on port 6000")
	log.Println("Connect using: telnet localhost 6000")

	// Start the server
	if err := srv.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
