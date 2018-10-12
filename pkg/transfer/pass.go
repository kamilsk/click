package transfer

import (
	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/service"
)

// PassRequest represents `GET /pass?url={URI}` request.
type PassRequest struct {
	Option service.Option
	Event  domain.RedirectEvent
}

// PassResponse represents `GET /pass?url={URI}` response.
type PassResponse struct {
	Error error
}
