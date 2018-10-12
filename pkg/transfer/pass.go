package transfer

import "github.com/kamilsk/click/pkg/domain"

// PassRequest represents `GET /pass?url={URL}` request.
type PassRequest struct {
	Context domain.RedirectContext
}

// PassResponse represents `GET /pass?url={URL}` response.
type PassResponse struct {
	Error      error
	StatusCode int
	URL        string
}
