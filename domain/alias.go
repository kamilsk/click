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

// Aliases holds set of Alias of the Link and provides useful methods for the set.
type Aliases []Alias

// Find tries to find a suitable alias by provided namespace and URN
// or returns an empty Alias if nothing found.
func (set Aliases) Find(namespace, urn string) Alias {
	var result Alias
	for _, alias := range set {
		if alias.Namespace == namespace && alias.URN == urn {
			return alias
		}
	}
	return result
}
