package transfer

import "go.octolab.org/ecosystem/click/internal/domain"

// PassRequest represents `GET /pass?url={URL}` request.
type PassRequest struct {
	Context domain.RedirectContext
}

// PassResponse represents `GET /pass?url={URL}` response.
type PassResponse struct {
	Error error
	URL   string
}
