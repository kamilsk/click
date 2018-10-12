package transfer

import "github.com/kamilsk/click/pkg/domain"

// PassRequest represents `GET /pass?url={URL}` request.
type PassRequest struct {
	Event domain.RedirectEvent
}

// PassResponse represents `GET /pass?url={URL}` response.
type PassResponse struct {
	Error      error
	StatusCode int
	URL        string
}
