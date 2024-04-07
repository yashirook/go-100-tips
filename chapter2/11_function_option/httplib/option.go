package httplib

import (
	"errors"
)

type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func NewServer(addr string, opts ...Option) (Server, error) {
	var options options
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return Server{}, err
		}
	}

	determinedPort, err := determinePort(options.port)
	if err != nil {
		return Server{}, err
	}

	// 本来はhttp.Serverを返すが今回のスコープ外
	return Server{Address: addr, Port: determinedPort}, nil
}

type Server struct {
	Address string
	Port    int
}
