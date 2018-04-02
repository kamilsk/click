package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kamilsk/click/domain"
	"github.com/kamilsk/click/errors"
	"github.com/kamilsk/click/server/middleware"
	"github.com/kamilsk/click/transfer"
	"github.com/kamilsk/click/transfer/api/v1"
)

// New returns a new instance of Click! server.
func New(service Service) *Server {
	return &Server{service: service}
}

// Server handles HTTP requests.
type Server struct {
	service Service
}

// GetV1 is responsible for `GET /api/v1/{UUID}` request handling.
func (s *Server) GetV1(rw http.ResponseWriter, req *http.Request) {
	var (
		id = req.Context().Value(middleware.LinkKey{}).(domain.UUID)
	)
	response := s.service.HandleGetV1(v1.GetRequest{ID: id})
	if response.Error != nil {
		if err, is := response.Error.(errors.ApplicationError); is {
			if _, is := err.IsClientError(); is {
				rw.WriteHeader(http.StatusNotFound)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response.Link)
}

// Redirect is responsible for `GET /{Alias}` request handling.
func (s *Server) Redirect(rw http.ResponseWriter, req *http.Request) {
	var (
		ns = fallback(req.Header.Get("X-Click-Namespace"), "global")
	)
	response := s.service.HandleRedirect(transfer.RedirectRequest{
		Namespace: ns,
		URN:       strings.Trim(req.URL.Path, "/"),
	})
	if response.Error != nil {
		if err, is := response.Error.(errors.ApplicationError); is {
			if _, is := err.IsClientError(); is {
				rw.WriteHeader(http.StatusNotFound)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	var statusCode int

	{ // domain logic
		switch {
		case response.Target.URI == "":
			statusCode = http.StatusNotImplemented
		case response.Alias.DeletedAt.Valid:
			statusCode = http.StatusMovedPermanently
		default:
			statusCode = http.StatusFound
		}
	}

	rw.Header().Set("Location", response.Target.URI)
	rw.WriteHeader(statusCode)
}

func fallback(value string, fallbackValues ...string) string {
	if value == "" {
		for _, value := range fallbackValues {
			if value != "" {
				return value
			}
		}
	}
	return value
}
