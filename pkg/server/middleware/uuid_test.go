package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/click/pkg/server/middleware"
)

const UUID domain.ID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestLink(t *testing.T) {
	tests := []struct {
		name string
		uuid domain.ID
		next func(uuid domain.ID) (*domain.ID, http.Handler)
		code int
	}{
		{"invalid uuid", "abc-def-ghi", func(uuid domain.ID) (*domain.ID, http.Handler) {
			return &uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			})
		}, http.StatusBadRequest},
		{"valid uuid", UUID, func(_ domain.ID) (*domain.ID, http.Handler) {
			uuid := new(domain.ID)
			return uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*uuid = req.Context().Value(LinkKey{}).(domain.ID)
			})
		}, http.StatusOK},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rw, req := httptest.NewRecorder(), &http.Request{}
			uuid, next := tc.next(tc.uuid)
			Link(tc.uuid.String(), rw, req, next)

			assert.Equal(t, tc.code, rw.Code)
			assert.Equal(t, tc.uuid, *uuid)
		})
	}
}
