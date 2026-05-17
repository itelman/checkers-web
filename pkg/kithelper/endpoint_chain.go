package kithelper

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func ChainMiddlewares(middlewares []endpoint.Middleware, others ...endpoint.Middleware) endpoint.Middleware {
	result := append(middlewares, others...)

	if len(middlewares) == 0 {
		return func(next endpoint.Endpoint) endpoint.Endpoint {
			return next
		}
	}

	return endpoint.Chain(result[0], result[1:]...)
}

type contextKey string

// operationNameContextKey holds the key used to store an operation name in the context.
const operationNameContextKey contextKey = "operationName"

// OperationNameMiddleware populates the context with a common name for the endpoint.
// It can be used in subsequent endpoints in the chain to identify the operation.
func OperationNameMiddleware(name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = context.WithValue(ctx, operationNameContextKey, name)

			return next(ctx, req)
		}
	}
}

// OperationName fetches the endpoint operation name from the context (if any).
// If an endpoint name is not found, or it isn't string, the second return argument is false.
func OperationName(ctx context.Context) (string, bool) {
	name, ok := ctx.Value(operationNameContextKey).(string)

	return name, ok
}
