package domain

import "database/sql"

// Alias represents an alias of the Link.
//go:generate easyjson -all
type Alias struct {
	ID        uint64         `json:"id"`
	LinkID    string         `json:"-"`
	Namespace string         `json:"namespace"`
	URN       string         `json:"urn"`
	CreatedAt string         `json:"-"`
	DeletedAt sql.NullString `json:"-"`
}
