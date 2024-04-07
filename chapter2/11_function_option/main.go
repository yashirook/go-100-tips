package main

import (
	"fmt"
)

func main() {
	fmt.Println("function option pattern")
}

type Config struct {
	Port int
}

type Server struct{}

// Configパターン
func NewServer(addr string, cfg Config) Server {
	return Server{}
}
