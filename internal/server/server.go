package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.octolab.org/ecosystem/click/internal/config"
	"go.octolab.org/ecosystem/click/internal/domain"
	"go.octolab.org/ecosystem/click/internal/errors"
	"go.octolab.org/ecosystem/click/internal/transfer"
	v1 "go.octolab.org/ecosystem/click/internal/transfer/api/v1"
)

// New returns a new instance of Click! server.
func New(cnf config.ServerConfig, service Service) *Server {
	return &Server{cnf, service}
}

// Server handles HTTP requests.
type Server struct {
	config  config.ServerConfig
	service Service
}

// GetV1 is responsible for `GET /api/v1/{Link.ID}` request handling.
// Deprecated: TODO issue#version3.0 use LinkEditor and gRPC gateway instead
func (s *Server) GetV1(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id := domain.ID(req.Form.Get("id"))
	if !id.IsValid() {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	resp := s.service.HandleGetV1(req.Context(), v1.GetRequest{ID: id})
	if resp.Error != nil {
		if err, is := resp.Error.(errors.ApplicationError); is {
			if _, is = err.IsClientError(); is {
				rw.WriteHeader(http.StatusNotFound)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(resp.Link)
}

// Pass is responsible for `GET /pass?url={URL}` request handling.
func (s *Server) Pass(rw http.ResponseWriter, req *http.Request) {
	resp := s.service.HandlePass(req.Context(), transfer.PassRequest{
		Context: domain.RedirectContext{
			Cookies: domain.FromCookies(req.Cookies()),
			Headers: domain.FromHeaders(req.Header),
			Queries: domain.FromRequest(req),
		},
	})
	if resp.Error != nil {
		if err, is := resp.Error.(errors.ApplicationError); is {
			if _, is = err.IsClientError(); is {
				rw.WriteHeader(http.StatusNotFound)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Location", resp.URL)
	rw.WriteHeader(http.StatusFound)
}

// Redirect is responsible for `GET /{Alias.URN}` request handling.
func (s *Server) Redirect(rw http.ResponseWriter, req *http.Request) {
	resp := s.service.HandleRedirect(req.Context(), transfer.RedirectRequest{
		Context: domain.RedirectContext{
			Cookies: domain.FromCookies(req.Cookies()),
			Headers: domain.FromHeaders(req.Header),
			Queries: domain.FromRequest(req),
		},
		URN: strings.Trim(req.URL.Path, "/"),
	})
	if resp.Error != nil {
		if err, is := resp.Error.(errors.ApplicationError); is {
			if _, is = err.IsClientError(); is {
				rw.WriteHeader(http.StatusNotFound)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Location", resp.URL)
	rw.WriteHeader(http.StatusFound)
}
