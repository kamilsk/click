package transfer

import "github.com/kamilsk/click/pkg/domain"

// RedirectRequest represents `GET /{Alias.URN}` request.
type RedirectRequest struct {
	Event domain.RedirectEvent
	URN   string
}

// RedirectResponse represents `GET /{Alias.URN}` response.
type RedirectResponse struct {
	Error      error
	StatusCode int
	URL        string
}
