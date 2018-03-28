package router

import "net/http"

// Server defines the behavior of Click! server.
type Server interface {
	// GetV1 is responsible for `GET /api/v1/{UUID}` request handling.
	GetV1(http.ResponseWriter, *http.Request)
	// Redirect is responsible for `GET /{Alias}` request handling.
	Redirect(http.ResponseWriter, *http.Request)
}
