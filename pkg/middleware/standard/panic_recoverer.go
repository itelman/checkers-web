package standard

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/itelman/checkers-web/pkg/kithelper"
	"github.com/ory/herodot"
)

func (m *middleware) PanicRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				m.logger.Error("PANIC RECOVERED",
					"error", err,
					"stack", string(debug.Stack()),
				)

				kithelper.WriteErrorCode(w, http.StatusInternalServerError, herodot.ErrInternalServerError.WithReason(fmt.Sprintf("%v", err)))
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
