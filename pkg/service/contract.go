package service

import (
	"context"

	"github.com/kamilsk/click/pkg/domain"
)

// Storage TODO issue#131
type Storage interface {
	// Link TODO issue#131
	Link(context.Context, domain.ID) (domain.Link, error)
	// LinkByAlias TODO issue#131
	LinkByAlias(ctx context.Context, ns domain.ID, urn string) (domain.Link, error)
}

// RedirectHandler TODO issue#131
type RedirectHandler interface {
	// LogRedirect stores a redirect event.
	LogRedirect(context.Context, domain.Redirect) error
}
