package service

import (
	"context"
	"net/http"

	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/transfer"
	"github.com/kamilsk/click/pkg/transfer/api/v1"
)

// New returns a new instance of Click! service.
func New(storage Storage, handler Tracker) *Click {
	return &Click{storage, handler}
}

// Click is the primary application service.
type Click struct {
	storage Storage
	tracker Tracker
}

// HandleGetV1 handles an input request.
// Deprecated: TODO issue#version3.0 use LinkEditor and gRPC gateway instead
func (service *Click) HandleGetV1(ctx context.Context, req v1.GetRequest) (resp v1.GetResponse) {
	resp.Link, resp.Error = service.storage.Link(ctx, req.ID)
	return
}

// HandlePass handles an input request.
func (service *Click) HandlePass(ctx context.Context, req transfer.PassRequest) (resp transfer.PassResponse) {
	req.Event.NamespaceID = req.Event.Context.Namespace()
	if !req.Event.NamespaceID.IsValid() {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("namespace is invalid"),
			"request %+v", req)
		return
	}

	req.Event.Code, req.Event.URL = http.StatusFound, req.Event.Context.URL()

	option := req.Event.Context.Option()
	if !option.NoLog {
		// if option.Anonymously {}
		req.Event.Identifier = "10000000-2000-4000-8000-160000000000" // TODO issue#134
		_ = service.tracker.LogRedirect(ctx, req.Event)               // TODO issue#51
	}

	return transfer.PassResponse{
		StatusCode: req.Event.Code,
		URL:        req.Event.URL,
	}
}

// HandleRedirect handles an input request.
func (service *Click) HandleRedirect(ctx context.Context, req transfer.RedirectRequest) (resp transfer.RedirectResponse) {
	req.Event.NamespaceID = req.Event.Context.Namespace()
	if !req.Event.NamespaceID.IsValid() {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("namespace is invalid"),
			"request %+v", req)
		return
	}

	link, err := service.storage.LinkByAlias(ctx, req.Event.NamespaceID, req.URN)
	if err != nil {
		resp.Error = err
		return
	}
	alias, found := link.Aliases.Find(req.Event.NamespaceID, req.URN)
	if !found {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("required alias not found"),
			"request %+v", req)
		return
	}
	target, found := link.Targets.Find(alias, req.Event.Context.Query)
	if !found {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("required target not found"),
			"request %+v", req)
		return
	}

	// if link.Deleted { http.StatusMovedPermanently ? }
	req.Event.Code, req.Event.URL = http.StatusFound, target.URL

	option := req.Event.Context.Option()
	if !option.NoLog {
		// if option.Anonymously {}
		req.Event.Identifier = "10000000-2000-4000-8000-160000000000" // TODO issue#134
		_ = service.tracker.LogRedirect(ctx, req.Event)               // TODO issue#51
	}

	return transfer.RedirectResponse{
		StatusCode: req.Event.Code,
		URL:        req.Event.URL,
	}
}
