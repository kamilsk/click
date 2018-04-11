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

const (
	globalNS    = "global"
	globalNSKey = "X-Click-Namespace"
	tokenKey    = "token"
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
		ns = fallback(req.Header.Get(globalNSKey), globalNS)
	)
	cookie, err := req.Cookie(tokenKey)
	if err != nil {
		cookie = &http.Cookie{Name: tokenKey}
	}

	response := s.service.HandleRedirect(transfer.RedirectRequest{
		EncryptedMarker: cookie.Value,
		Namespace:       ns,
		URN:             strings.Trim(req.URL.Path, "/"),
		Query:           req.URL.Query(),
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

	cookie.MaxAge, cookie.Path, cookie.Value = 0, "/", response.EncryptedMarker
	cookie.Secure, cookie.HttpOnly = true, true
	http.SetCookie(rw, cookie)
	rw.Header().Set("Location", response.Target.URI)
	rw.WriteHeader(statusCode)

	go func() {
		cookie := make(map[string]string, len(req.Cookies())+1)
		for _, c := range req.Cookies() {
			if c.HttpOnly {
				cookie[c.Name] = c.Value
			}
		}
		cookie[tokenKey] = response.EncryptedMarker
		header := make(map[string][]string, len(req.Header))
		for key, values := range req.Header {
			switch {
			case key == "Accept":
				continue
			case key == "Cookie":
				continue
			default:
				header[key] = values
			}
		}
		s.service.LogRedirectEvent(domain.Log{
			LinkID:   response.Alias.LinkID,
			AliasID:  response.Alias.ID,
			TargetID: response.Target.ID,
			URI:      response.Target.URI,
			Code:     statusCode,
			Context:  domain.Metadata{Cookie: cookie, Header: header},
		})
	}()
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
