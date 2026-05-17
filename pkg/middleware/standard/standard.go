package standard

import (
	"log/slog"
)

type middleware struct {
	logger *slog.Logger
}

func NewMiddleware(logger *slog.Logger) *middleware {
	return &middleware{
		logger: logger,
	}
}
