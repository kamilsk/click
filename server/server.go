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
	globalNS        = "global"
	namespaceHeader = "X-Click-Namespace"
	optionsHeader   = "X-Click-Options"
	passQueryParam  = "url"
	tokenCookieName = "token"
)

// New returns a new instance of Click! server.
func New(service Service) *Server {
	return &Server{service: service}
}

// Server handles HTTP requests.
type Server struct {
	service Service
}

// GetV1 is responsible for `GET /api/v1/{Link.ID}` request handling.
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

// Pass is responsible for `GET /pass?url={URI}` request handling.
func (s *Server) Pass(rw http.ResponseWriter, req *http.Request) {
	to := req.URL.Query().Get(passQueryParam)
	if to == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var opts options
	opts.From(req.Header.Get(optionsHeader))

	// TODO: move to middleware layer
	// TODO: support opts.Anonymously()
	cookie, err := req.Cookie(tokenCookieName)
	if err != nil {
		cookie = &http.Cookie{Name: tokenCookieName}
	}

	response := s.service.HandlePass(transfer.PassRequest{EncryptedMarker: cookie.Value})
	if response.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: move to middleware layer
	// TODO: support opts.Anonymously()
	cookie.MaxAge, cookie.Path, cookie.Value = 0, "/", response.EncryptedMarker
	cookie.Secure, cookie.HttpOnly = true, true
	http.SetCookie(rw, cookie)

	rw.Header().Set("Location", to)
	rw.WriteHeader(http.StatusFound)

	if !opts.NoLog() {
		go log(s.service.LogRedirectEvent, req, response.EncryptedMarker, domain.Log{
			LinkID:   string(domain.EmptyUUID),
			AliasID:  0,
			TargetID: 0,
			URI:      to,
			Code:     http.StatusFound,
		})
	}
}

// Redirect is responsible for `GET /{Alias.URN}` request handling.
func (s *Server) Redirect(rw http.ResponseWriter, req *http.Request) {
	var ns = fallback(req.Header.Get(namespaceHeader), globalNS)

	var opts options
	opts.From(req.Header.Get(optionsHeader))

	// TODO: move to middleware layer
	// TODO: support opts.Anonymously()
	cookie, err := req.Cookie(tokenCookieName)
	if err != nil {
		cookie = &http.Cookie{Name: tokenCookieName}
	}

	response := s.service.HandleRedirect(transfer.RedirectRequest{
		EncryptedMarker: cookie.Value,
		Namespace:       ns,
		URN:             strings.Trim(req.URL.Path, "/"),
		Query:           req.URL.Query(),
	})

	// TODO: move to middleware layer
	// TODO: support opts.Anonymously()
	cookie.MaxAge, cookie.Path, cookie.Value = 0, "/", response.EncryptedMarker
	cookie.Secure, cookie.HttpOnly = true, true
	http.SetCookie(rw, cookie)

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

	if !opts.NoLog() {
		go log(s.service.LogRedirectEvent, req, response.EncryptedMarker, domain.Log{
			LinkID:   response.Alias.LinkID,
			AliasID:  response.Alias.ID,
			TargetID: response.Target.ID,
			URI:      response.Target.URI,
			Code:     statusCode,
		})
	}
}

type options []string

func (opts *options) From(str string) {
	s := strings.Split(str, ";")
	*opts = make([]string, 0, len(s))
	for _, str := range s {
		*opts = append(*opts, strings.ToLower(strings.TrimSpace(str)))
	}
}

func (opts options) Anonymously() bool {
	return opts.find("anonym")
}

func (opts options) Debug() bool {
	return opts.find("debug")
}

func (opts options) NoLog() bool {
	return opts.find("nolog")
}

func (opts options) find(str string) bool {
	for _, opt := range opts {
		if opt == str {
			return true
		}
	}
	return false
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

func log(handle func(event domain.Log), req *http.Request, token string, event domain.Log) {
	var (
		cookie map[string]string
		header map[string][]string
		query  map[string][]string
	)
	{
		origin := req.Cookies()
		cookie = make(map[string]string, len(origin))
		for _, c := range origin {
			if c.HttpOnly && c.Secure {
				cookie[c.Name] = c.Value
			}
		}
		cookie[tokenCookieName] = token
	}
	{
		origin := req.Header
		header = make(map[string][]string, len(origin))
		for key, values := range origin {
			if key != "Cookie" {
				header[key] = values
			}
		}
	}
	{
		origin := req.URL.Query()
		query = make(map[string][]string, len(origin))
		for key, values := range origin {
			if key != passQueryParam {
				query[key] = values
			}
		}
	}
	event.Context = domain.Metadata{Cookie: cookie, Header: header, Query: query}
	handle(event)
}
