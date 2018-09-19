package query

import "github.com/kamilsk/click/pkg/domain"

// CreateTarget TODO issue#131
type CreateTarget struct {
	ID *domain.ID
}

// ReadTarget TODO issue#131
type ReadTarget struct {
	ID domain.ID
}

// UpdateTarget TODO issue#131
type UpdateTarget struct {
	ID domain.ID
}

// DeleteTarget TODO issue#131
type DeleteTarget struct {
	ID          domain.ID
	Permanently bool
}
