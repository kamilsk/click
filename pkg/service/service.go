package service

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/kamilsk/click/pkg/config"
	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/transfer"
	"github.com/kamilsk/click/pkg/transfer/api/v1"
)

// New returns a new instance of Click! service.
func New(cnf config.ServiceConfig, storage Storage, tracker Tracker) *Click {
	return &Click{cnf, storage, tracker}
}

// Click is the primary application service.
type Click struct {
	config  config.ServiceConfig
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

	resp.URL = req.Context.Redirect()
	if resp.URL == "" {
		resp.Error = errors.NotFound(errors.LinkNotFoundMessage, errors.Simple("url is empty"),
			"request %+v", req)
		return
	}

	// issue#123, try to decode encoded url
	if !strings.HasPrefix(resp.URL, "http") {
		// url=aHR0cHM6Ly9naXRodWIuY29tL2thbWlsc2svY2xpY2sK -> url=https://github.com/kamilsk/click
		if raw, decodeErr := base64.URLEncoding.DecodeString(resp.URL); decodeErr == nil {
			resp.URL = string(raw)
		}
	}

	if !req.Context.Option().NoLog && !ignore(req.Context) {
		// if option.Anonymously {}
		event := domain.RedirectEvent{
			NamespaceID: ns,
			Identifier:  nil, // TODO issue#134
			Context:     req.Context,
			Code:        http.StatusFound, // TODO issue#design
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
	resp.URL = target.URL

	if !req.Context.Option().NoLog && !ignore(req.Context) {
		// if option.Anonymously {}
		event := domain.RedirectEvent{
			NamespaceID: ns,
			LinkID:      &link.ID,
			AliasID:     &alias.ID,
			TargetID:    &target.ID,
			Identifier:  nil, // TODO issue#134
			Context:     req.Context,
			Code:        http.StatusFound, // TODO issue#design
			URL:         resp.URL,
		}
		resp.Error = service.tracker.LogRedirect(ctx, event) // TODO issue#51
	}

	return
}

// TODO issue#refactoring check and log async, add blacklist
// - https://github.com/monperrus/crawler-user-agents/blob/master/crawler-user-agents.json
func ignore(req domain.RedirectContext) bool {
	// issue#63 do not log curl request
	return strings.HasPrefix(req.UserAgent(), "curl/")
}
