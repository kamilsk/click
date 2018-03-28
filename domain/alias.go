package domain

import "database/sql"

// Alias represents an alias of the Link.
type Alias struct {
	ID        int64          `json:"id"`
	LinkID    string         `json:"-"`
	Namespace string         `json:"namespace"`
	URN       string         `json:"urn"`
	CreatedAt string         `json:"-"`
	DeletedAt sql.NullString `json:"-"`
}
