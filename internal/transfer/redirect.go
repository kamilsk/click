package transfer

import "go.octolab.org/ecosystem/click/internal/domain"

// RedirectRequest represents `GET /{Alias.URN}` request.
type RedirectRequest struct {
	Context domain.RedirectContext
	URN     string
}

// RedirectResponse represents `GET /{Alias.URN}` response.
type RedirectResponse struct {
	Error error
	URL   string
}
