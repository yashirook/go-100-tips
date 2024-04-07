package main

import (
	"fmt"
	"log"
	"option/httplib"
)

func main() {
	// # builder pattern
	// builder := httplib.ConfigBuilder{}
	// builder.Port(8080)
	// cfg, err := builder.Build()
	// if err != nil {
	// 	log.Fatalf("Failed to build config: %v", err)
	// }
	// server := NewServer("localhost", cfg)

	// # function option pattern
	server, err := httplib.NewServer("localhost", httplib.WithPort(8080))
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	fmt.Printf("server is successfully initialized with port: %d\n", server.Port)
}

type Server struct {
	Port int
}

// Configパターン
func NewServer(addr string, cfg httplib.Config) Server {
	return Server{
		Port: cfg.Port,
	}
}
