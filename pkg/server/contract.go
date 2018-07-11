package server

import (
	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/transfer"
	"github.com/kamilsk/click/pkg/transfer/api/v1"
)

// Service defines the behavior of Click! service.
type Service interface {
	// HandleGetV1 handles an input request.
	HandleGetV1(v1.GetRequest) v1.GetResponse
	// HandlePass handles an input request.
	HandlePass(transfer.PassRequest) transfer.PassResponse
	// HandleRedirect handles an input request.
	HandleRedirect(transfer.RedirectRequest) transfer.RedirectResponse

	// LogRedirectEvent stores a "redirect event".
	LogRedirectEvent(event domain.Log)
}
