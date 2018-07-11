package domain

import "database/sql"

// Link represents a "redirect entity".
//go:generate easyjson -all
type Link struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Status    string         `json:"status"`
	CreatedAt string         `json:"-"`
	UpdatedAt sql.NullString `json:"-"`
	Aliases   Aliases        `json:"aliases"`
	Targets   Targets        `json:"targets"`
}
