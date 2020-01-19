package main

import (
	"fmt"
	"log"
	"neo4j/internal/di"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:6565")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server, _, _ := di.InitApp()
	fmt.Println("server starting")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	fmt.Println("server started")
}
