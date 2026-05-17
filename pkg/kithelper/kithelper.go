package kithelper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/itelman/checkers-web/pkg/jsonhelper"
	"github.com/ory/herodot"
)

type encoder struct {
	logger *slog.Logger
}

func NewEncoder(logger *slog.Logger) *encoder {
	return &encoder{logger: logger}
}

func WriteErrorInternalServer(logger *slog.Logger, w http.ResponseWriter, err error) {
	logger.Error("INTERNAL SERVER ERROR",
		"error", err,
		"stack", string(debug.Stack()),
	)

	WriteErrorCode(w, http.StatusInternalServerError, herodot.ErrInternalServerError.WithReason(err.Error()))
	return
}

type ErrorJSON interface {
	JSON() any
}

func (e *encoder) ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	c := herodot.StatusCodeCarrier(nil)

	if errors.As(err, &c) {
		WriteErrorCode(w, c.StatusCode(), err)
	} else {
		WriteErrorInternalServer(e.logger, w, err)
	}
}

type errorWithStack interface {
	Cause() error
}

func errorIsWithStack(err error) bool {
	_, ok := err.(errorWithStack)
	return ok
}

func WriteErrorCode(w http.ResponseWriter, code int, err error) {
	var payload interface{} = err

	if errorIsWithStack(err) {
		payload = errors.Unwrap(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func EmptyResponse(_ context.Context, w http.ResponseWriter, _ interface{}) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func EncodeResponse[A, B comparable](fn jsonhelper.Encoder[A, B]) kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		output, ok := response.(A)
		if !ok {
			return http.ErrNotSupported
		}
		jsonOutput := fn(output)

		return kithttp.EncodeJSONResponse(ctx, w, jsonOutput)
	}
}

func EncodeRedirectResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	url, ok := response.(string)
	if !ok {
		return http.ErrNotSupported
	}

	http.Redirect(w, &http.Request{}, url, http.StatusFound)
	return nil
}

func EmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil //nolint:nilnil  // uneccessary error check.
}

var ErrorCastFailed = fmt.Errorf("[KITHELPER ERROR]: %s", "cast failed")
