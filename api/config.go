package main

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"
)

type config struct {
	ENV       string `env:"ENV, default=prod"`
	Port      string `env:"PORT, default=8888"`
	APIHost   string `env:"API_HOST, default=http://localhost:8888"`
	JWTSecret string `env:"JWT_SECRET"`
}

func newConfig(ctx context.Context) (*config, error) {
	var c config
	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
