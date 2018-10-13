package service

import (
	"context"
	"net/http"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/transfer"
	"github.com/kamilsk/click/pkg/transfer/api/v1"
)

// New returns a new instance of Click! service.
func New(storage Storage, tracker Tracker) *Click {
	return &Click{storage, tracker}
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
	ns := req.Context.Namespace()
	if !ns.IsValid() {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("namespace is invalid"),
			"request %+v", req)
		return
	}

	resp.StatusCode, resp.URL = http.StatusFound, req.Context.Redirect()
	if resp.URL == "" {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("url is empty"),
			"request %+v", req)
		return
	}

	option := req.Context.Option()
	if !option.NoLog {
		// if option.Anonymously {}
		event := domain.RedirectEvent{
			NamespaceID: ns,
			Identifier:  nil, // TODO issue#134
			Context:     req.Context,
			Code:        resp.StatusCode,
			URL:         resp.URL,
		}
		_ = service.tracker.LogRedirect(ctx, event) // TODO issue#51
	}

	return
}

// HandleRedirect handles an input request.
func (service *Click) HandleRedirect(ctx context.Context, req transfer.RedirectRequest) (resp transfer.RedirectResponse) {
	ns := req.Context.Namespace()
	if !ns.IsValid() {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("namespace is invalid"),
			"request %+v", req)
		return
	}

	link, err := service.storage.LinkByAlias(ctx, ns, req.URN)
	if err != nil {
		resp.Error = err
		return
	}
	alias, found := link.Aliases.Find(ns, req.URN)
	if !found {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("required alias not found"),
			"request %+v", req)
		return
	}
	target, found := link.Targets.Find(alias, req.Context.Queries)
	if !found {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("required target not found"),
			"request %+v", req)
		return
	}

	// if link.Deleted { http.StatusMovedPermanently ? }
	resp.StatusCode, resp.URL = http.StatusFound, target.URL

	option := req.Context.Option()
	if !option.NoLog {
		// if option.Anonymously {}
		event := domain.RedirectEvent{
			NamespaceID: ns,
			LinkID:      &link.ID,
			AliasID:     &alias.ID,
			TargetID:    &target.ID,
			Identifier:  nil, // TODO issue#134
			Context:     req.Context,
			Code:        resp.StatusCode,
			URL:         resp.URL,
		}
		resp.Error = service.tracker.LogRedirect(ctx, event) // TODO issue#51
	}

	return
}
