package httplib

import (
	"errors"
	"fmt"
	"net/http"
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

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}

	determinedPort, err := determinePort(options.port)
	if err != nil {
		return nil, err
	}

	fmt.Printf("initialize http.Server struct using by %d", determinedPort)
	// 本来はhttp.Serverを返すが今回のスコープ外
	return nil, nil
}
