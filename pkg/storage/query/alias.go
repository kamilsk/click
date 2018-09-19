package query

import "github.com/kamilsk/click/pkg/domain"

// CreateAlias TODO issue#131
type CreateAlias struct {
	ID *domain.ID
}

// ReadAlias TODO issue#131
type ReadAlias struct {
	ID domain.ID
}

// UpdateAlias TODO issue#131
type UpdateAlias struct {
	ID domain.ID
}

// DeleteAlias TODO issue#131
type DeleteAlias struct {
	ID          domain.ID
	Permanently bool
}
