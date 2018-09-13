package middleware

import (
	"context"
	"net/http"

	"github.com/kamilsk/click/pkg/domain"
)

// Link validates the passed Link ID and injects it to the request context.
func Link(uuid string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	if !domain.ID(uuid).IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), LinkKey{}, domain.ID(uuid))))
}
