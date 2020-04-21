package query

import "go.octolab.org/ecosystem/click/internal/domain"

// CreateAlias TODO issue#131
type CreateAlias struct {
	ID          *domain.ID
	LinkID      domain.ID
	NamespaceID domain.ID
	URN         string
}

// ReadAlias TODO issue#131
type ReadAlias struct {
	ID domain.ID
}

// UpdateAlias TODO issue#131
type UpdateAlias struct {
	ID          domain.ID
	LinkID      domain.ID
	NamespaceID domain.ID
	URN         string
}

// DeleteAlias TODO issue#131
type DeleteAlias struct {
	ID          domain.ID
	Permanently bool
}
