package dynamic

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func RequireUserMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			/*
				if auth.GetCurrentUser(ctx) == nil {
					return nil, herodot.ErrUnauthorized
				}

			*/

			return next(ctx, request)
		}
	}
}
