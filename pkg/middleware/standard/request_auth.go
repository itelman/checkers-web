package standard

import (
	"net/http"
	"strings"
)

func (m *middleware) RequestAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerVal := r.Header.Get("Authorization")
		if headerVal == "" || !strings.HasPrefix(headerVal, "Bearer ") {
			next.ServeHTTP(w, r)
			return
		}

		/*
			c := herodot.StatusCodeCarrier(nil)

			token := strings.TrimPrefix(headerVal, "Bearer ")
			resp, err := m.service.AuthenticateRequest(r.Context(), auth.AuthenticateInput{Token: token})
			if errors.As(err, &c) {
				kithelper.WriteErrorCode(w, c.StatusCode(), err)
				return
			} else if err != nil {
				kithelper.WriteErrorInternalServer(m.logger, w, err)
				return
			}

			ctx := context.WithValue(r.Context(), auth.ContextKeyUser, resp.User)
			next.ServeHTTP(w, r.WithContext(ctx))
		*/
	})
}
