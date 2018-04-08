package domain

import "database/sql"

// Target represents a target of the Link.
//go:generate easyjson -all
type Target struct {
	ID        uint64         `json:"id"`
	LinkID    string         `json:"-"`
	URI       string         `json:"uri"`
	Rule      Rule           `json:"rule"`
	CreatedAt string         `json:"-"`
	UpdatedAt sql.NullString `json:"-"`
}
