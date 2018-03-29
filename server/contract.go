package server

import (
	"github.com/kamilsk/click/transfer"
	"github.com/kamilsk/click/transfer/api/v1"
)

// Service defines the behavior of Click! service.
type Service interface {
	// HandleGetV1 handles an input request.
	HandleGetV1(v1.GetRequest) v1.GetResponse
	// HandleRedirect handles an input request.
	HandleRedirect(transfer.RedirectRequest) transfer.RedirectResponse
}
