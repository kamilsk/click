package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/server/middleware"
	"github.com/kamilsk/click/pkg/transfer"
	"github.com/kamilsk/click/pkg/transfer/api/v1"
)

// New returns a new instance of Click! server.
func New(service Service) *Server {
	return &Server{service}
}

// Server handles HTTP requests.
type Server struct {
	service Service
}

// GetV1 is responsible for `GET /api/v1/{Link.ID}` request handling.
// Deprecated: TODO issue#version3.0 use LinkEditor and gRPC gateway instead
func (s *Server) GetV1(rw http.ResponseWriter, req *http.Request) {
	var id = req.Context().Value(middleware.LinkKey{}).(domain.ID)
	response := s.service.HandleGetV1(req.Context(), v1.GetRequest{ID: id})
	if response.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response.Link)
}

// Pass is responsible for `GET /pass?url={URL}` request handling.
func (s *Server) Pass(rw http.ResponseWriter, req *http.Request) {
	resp := s.service.HandlePass(req.Context(), transfer.PassRequest{
		Context: domain.RedirectContext{
			Cookies: func() map[string]string {
				cookies := make(map[string]string)
				for _, cookie := range req.Cookies() {
					if cookie.HttpOnly && cookie.Secure {
						cookies[cookie.Name] = cookie.Value
					}
				}
				return cookies
			}(),
			Headers: func() map[string][]string {
				headers := make(map[string][]string)
				for key, values := range req.Header {
					if key != "Cookie" {
						headers[key] = values
					}
				}
				return headers
			}(),
			Queries: req.URL.Query(),
		},
	})
	if resp.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Location", resp.URL)
	rw.WriteHeader(resp.StatusCode)
}

// Redirect is responsible for `GET /{Alias.URN}` request handling.
func (s *Server) Redirect(rw http.ResponseWriter, req *http.Request) {
	resp := s.service.HandleRedirect(req.Context(), transfer.RedirectRequest{
		Context: domain.RedirectContext{
			Cookies: func() map[string]string {
				cookies := make(map[string]string)
				for _, cookie := range req.Cookies() {
					if cookie.HttpOnly && cookie.Secure {
						cookies[cookie.Name] = cookie.Value
					}
				}
				return cookies
			}(),
			Headers: func() map[string][]string {
				headers := make(map[string][]string)
				for key, values := range req.Header {
					if key != "Cookie" {
						headers[key] = values
					}
				}
				return headers
			}(),
			Queries: req.URL.Query(),
		},
		URN: strings.Trim(req.URL.Path, "/"),
	})
	if resp.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Location", resp.URL)
	rw.WriteHeader(resp.StatusCode)
}
