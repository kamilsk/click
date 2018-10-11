package service

import (
	"context"

	"github.com/kamilsk/click/pkg/transfer"
	"github.com/kamilsk/click/pkg/transfer/api/v1"
)

// New returns a new instance of Click! service.
func New(storage Storage, handler RedirectHandler) *Click {
	return &Click{storage, handler}
}

// Click is the primary application service.
type Click struct {
	storage Storage
	handler RedirectHandler
}

// HandleGetV1 handles an input request.
// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
func (service *Click) HandleGetV1(ctx context.Context, request v1.GetRequest) v1.GetResponse {
	var response v1.GetResponse
	response.Link, response.Error = service.storage.Link(ctx, request.ID)
	return response
}

// HandlePass handles an input request.
func (service *Click) HandlePass(ctx context.Context, request transfer.PassRequest) transfer.PassResponse {
	var response transfer.PassResponse

	// TODO issue#51
	if !request.Option.NoLog {
		// if request.Option.Anonymously {}
		request.Redirect.Identifier = "10000000-2000-4000-8000-160000000000" // TODO issue#134
		_ = service.handler.LogRedirect(ctx, request.Redirect)
	}

	return response
}

// HandleRedirect handles an input request.
func (service *Click) HandleRedirect(ctx context.Context, request transfer.RedirectRequest) transfer.RedirectResponse {
	var response transfer.RedirectResponse
	link, err := service.storage.LinkByAlias(ctx, request.Namespace, request.URN)
	if err != nil {
		response.Error = err
		return response
	}
	response.Alias = link.Aliases.Find(request.Namespace, request.URN)
	response.Target = link.Targets.Find(response.Alias, request.Query)

	// TODO issue#51
	if !request.Option.NoLog {
		// if request.Option.Anonymously {}
		request.Redirect.Identifier = "10000000-2000-4000-8000-160000000000" // TODO issue#134
		_ = service.handler.LogRedirect(ctx, request.Redirect)
	}

	return response
}
