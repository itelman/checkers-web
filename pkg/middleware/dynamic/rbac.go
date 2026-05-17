package dynamic

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

func RBACMiddleware(roles map[string]interface{}) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if len(roles) == 0 {
				return nil, fmt.Errorf("[RBAC ERROR]: no roles provided")
			}

			/*
				user := auth.GetCurrentUser(ctx)
				for _, role := range user.Roles {
					if _, ok := roles[role]; ok {
						return next(ctx, request)
					}
				}

				return nil, herodot.ErrForbidden.WithReason(fmt.Sprintf("Required role(s): %s, received role(s): %s", roles, user.Roles))

			*/

			return nil, nil
		}
	}
}
