package router

import "net/http"

// Server defines the behavior of Click! server.
type Server interface {
	// GetV1 is responsible for `GET /api/v1/{Link.ID}` request handling.
	GetV1(http.ResponseWriter, *http.Request)
	// Pass is responsible for `GET /pass?url={URI}` request handling.
	Pass(http.ResponseWriter, *http.Request)
	// Redirect is responsible for `GET /{Alias.URN}` request handling.
	Redirect(http.ResponseWriter, *http.Request)
}
