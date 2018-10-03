package query

import "github.com/kamilsk/click/pkg/domain"

// CreateLink TODO issue#131
type CreateLink struct {
	ID   *domain.ID
	Name string
}

// ReadLink TODO issue#131
type ReadLink struct {
	ID domain.ID
}

// UpdateLink TODO issue#131
type UpdateLink struct {
	ID   domain.ID
	Name string
}

// DeleteLink TODO issue#131
type DeleteLink struct {
	ID          domain.ID
	Permanently bool
}
