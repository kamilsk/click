//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package router_test -destination mock_server_test.go go.octolab.org/ecosystem/click/internal/server/router Server
package router_test

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/middleware"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/ecosystem/click/internal/server/router"
	"go.octolab.org/ecosystem/click/internal/server/router/chi"
)

const uuid = "10000000-2000-4000-8000-160000000005"

func TestMain(m *testing.M) {
	before := middleware.DefaultLogger
	middleware.DefaultLogger = func(next http.Handler) http.Handler { return next }
	defer func() { middleware.DefaultLogger = before }()
	os.Exit(m.Run())
}

func TestContract_chi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name string
		api  func() Server
		req  func() *http.Request
	}{
		{"check GetV1 method", func() Server {
			mock := NewMockServer(ctrl)
			mock.EXPECT().
				GetV1(gomock.Any(), gomock.Any()).
				Do(func(rw http.ResponseWriter, req *http.Request) {
					assert.Equal(t, uuid, req.FormValue("id"))
				})
			return mock
		}, func() *http.Request {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/%s", uuid), nil)
			return req
		}},
		{"check Pass method", func() Server {
			mock := NewMockServer(ctrl)
			mock.EXPECT().
				Pass(gomock.Any(), gomock.Any()).
				Do(func(rw http.ResponseWriter, req *http.Request) {
					assert.Equal(t, uuid, req.FormValue("url"))
				})
			return mock
		}, func() *http.Request {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/pass?url=%s", uuid), nil)
			return req
		}},
		{"check Redirect method", func() Server {
			mock := NewMockServer(ctrl)
			mock.EXPECT().
				Redirect(gomock.Any(), gomock.Any()).
				Do(func(rw http.ResponseWriter, req *http.Request) {
					assert.Equal(t, uuid, strings.Trim(req.URL.Path, "/"))
				})
			return mock
		}, func() *http.Request {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", uuid), nil)
			return req
		}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			r := chi.NewRouter(tc.api())
			r.ServeHTTP(nil, tc.req())
		})
	}
}
