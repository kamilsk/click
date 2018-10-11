package server

import (
	"context"

	"github.com/kamilsk/click/pkg/transfer"
	"github.com/kamilsk/click/pkg/transfer/api/v1"
)

// Service defines the behavior of Click! service.
type Service interface {
	// HandleGetV1 TODO issue#131
	HandleGetV1(context.Context, v1.GetRequest) v1.GetResponse
	// HandlePass TODO issue#131
	HandlePass(context.Context, transfer.PassRequest) transfer.PassResponse
	// HandleRedirect TODO issue#131
	HandleRedirect(context.Context, transfer.RedirectRequest) transfer.RedirectResponse
}
