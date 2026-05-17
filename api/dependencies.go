package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/itelman/checkers-web/pkg/jwthelper"
)

type dependencies struct {
	jwt    *jwthelper.Client
	logger *slog.Logger
}

func (d *dependencies) Close() {
	if d == nil {
		return
	}
}

type option func(context.Context, *dependencies) error

func NewDependencies(ctx context.Context, opts ...option) (deps *dependencies, err error) {
	defer func() {
		if err != nil {
			deps.Close()
		}
	}()

	deps = &dependencies{}
	for _, opt := range opts {
		if err := opt(ctx, deps); err != nil {
			return nil, err
		}
	}

	return deps, nil
}

func WithJWTClient(secret string) option {
	return func(_ context.Context, d *dependencies) error {
		d.jwt = jwthelper.NewClient(secret)
		return nil
	}
}

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

func WithLogger(env string) option {
	return func(_ context.Context, d *dependencies) error {
		var logLvl slog.Level
		switch env {
		case EnvLocal:
			logLvl = slog.LevelDebug
		case EnvDev:
			logLvl = slog.LevelDebug
		case EnvProd:
			logLvl = slog.LevelInfo
		default:
			logLvl = slog.LevelInfo
		}

		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLvl,
		}))
		slog.SetDefault(logger)
		d.logger = logger

		return nil
	}
}
