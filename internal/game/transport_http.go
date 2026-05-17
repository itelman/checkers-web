package game

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/itelman/checkers-web/pkg/kithelper"
	"github.com/ory/herodot"
)

func newGameHandler(endpoint endpoint.Endpoint, options ...kithttp.ServerOption) http.Handler {
	return kithttp.NewServer(
		endpoint,
		kithelper.EmptyRequest,
		kithttp.EncodeJSONResponse,
		options...,
	)
}

func validateMoveHandler(endpoint endpoint.Endpoint, options ...kithttp.ServerOption) http.Handler {
	return kithttp.NewServer(
		endpoint,
		decodeValidateMove,
		kithttp.EncodeJSONResponse,
		options...,
	)
}

func RegisterHTTPHandlers(router *mux.Router, endpoints Endpoints, options ...kithttp.ServerOption) {
	router.Methods(http.MethodPost).Path("/new").Handler(newGameHandler(endpoints.NewGameEndpoint, options...))
	router.Methods(http.MethodGet).Path("/move").Handler(validateMoveHandler(endpoints.ValidateMoveEndpoint, options...))
}

func decodeValidateMove(_ context.Context, r *http.Request) (interface{}, error) {
	var input ValidateMoveInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, herodot.ErrBadRequest.WithReason("Invalid request body")
	}
	defer r.Body.Close()

	return input, nil
}
