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

	if b.port == nil {
		cfg.Port = DefaultHTTPPort
	}

	if *b.port == 0 {
		cfg.Port = randomPort()
	} else if *b.port < 0 {
		return Config{}, errors.New("invalid port number: port number should be positive value.")
	} else {
		cfg.Port = *b.port
	}

	return cfg, nil
}

func randomPort() int {
	return rand.Int()
}

type Config struct {
	Port int
}
