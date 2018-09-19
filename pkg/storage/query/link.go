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

// WriteLog TODO issue#131
type WriteLog struct {
	LinkID          domain.ID
	AliasID         domain.ID
	TargetID        domain.ID
	Identifier      domain.ID
	URI             string
	Code            uint16
	RedirectContext domain.RedirectContext
}
