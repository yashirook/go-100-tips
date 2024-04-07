package main

import (
	"fmt"
	"log"
	"option/httplib"
)

func main() {
	fmt.Println("function option pattern")
	builder := httplib.ConfigBuilder{}
	builder.Port(8080)
	cfg, err := builder.Build()
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	server := NewServer("localhost", cfg)
	log.Printf("server: %v", server)
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
