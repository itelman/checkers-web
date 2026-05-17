package game

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	NewGameEndpoint      endpoint.Endpoint
	ValidateMoveEndpoint endpoint.Endpoint
}

type GameResponse struct {
	State GameState `json:"state"`
	Err   string    `json:"err,omitempty"`
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		NewGameEndpoint:      makeNewGameEndpoint(s),
		ValidateMoveEndpoint: makeValidateMoveEndpoint(s),
	}
}

func makeNewGameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		state, err := s.NewGame(ctx)
		if err != nil {
			return GameResponse{Err: err.Error()}, nil
		}
		return GameResponse{State: state}, nil
	}
}

func makeValidateMoveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		state, err := s.ValidateMove(ctx, request.(ValidateMoveInput))
		if err != nil {
			return GameResponse{Err: err.Error()}, nil
		}
		return GameResponse{State: state}, nil
	}
}
