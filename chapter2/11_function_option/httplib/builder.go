package httplib

import (
	"math/rand"

	"github.com/pkg/errors"
)

type ConfigBuilder struct {
	port *int
}

const DefaultHTTPPort = 80

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	determinedPort, err := determinePort(b.port)
	if err != nil {
		return Config{}, err
	}

	cfg.Port = determinedPort
	return cfg, nil
}

func randomPort() int {
	return rand.Int()
}

type Config struct {
	Port int
}

func determinePort(port *int) (int, error) {
	if port == nil {
		return DefaultHTTPPort, nil
	} else {
		if *port == 0 {
			return randomPort(), nil
		} else if *port < 0 {
			return 0, errors.New("invalid port number: port number should be positive value.")
		} else {
			return *port, nil
		}
	}
}
