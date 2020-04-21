package service

import (
	"context"

	"go.octolab.org/ecosystem/click/internal/domain"
)

// Storage TODO issue#131
type Storage interface {
	// Link TODO issue#131
	Link(context.Context, domain.ID) (domain.Link, error)
	// LinkByAlias TODO issue#131
	LinkByAlias(ctx context.Context, ns domain.ID, urn string) (domain.Link, error)
}

// Tracker TODO issue#131
type Tracker interface {
	// LogRedirect stores a redirect event.
	LogRedirect(context.Context, domain.RedirectEvent) error
}
