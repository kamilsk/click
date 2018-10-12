package transfer

import "github.com/kamilsk/click/pkg/domain"

// RedirectRequest represents `GET /{Alias.URN}` request.
type RedirectRequest struct {
	Context domain.RedirectContext
	URN     string
}

// RedirectResponse represents `GET /{Alias.URN}` response.
type RedirectResponse struct {
	Error      error
	StatusCode int
	URL        string
}
