package server

import (
	"context"

	"go.octolab.org/ecosystem/click/internal/transfer"
	v1 "go.octolab.org/ecosystem/click/internal/transfer/api/v1"
)

// Service defines the behavior of Click! service.
type Service interface {
	// HandleGetV1 TODO issue#131
	// Deprecated: TODO issue#version3.0 use LinkEditor and gRPC gateway instead
	HandleGetV1(context.Context, v1.GetRequest) v1.GetResponse
	// HandlePass TODO issue#131
	HandlePass(context.Context, transfer.PassRequest) transfer.PassResponse
	// HandleRedirect TODO issue#131
	HandleRedirect(context.Context, transfer.RedirectRequest) transfer.RedirectResponse
}
