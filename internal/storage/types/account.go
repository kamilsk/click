package types

import (
	"time"

	"go.octolab.org/ecosystem/click/internal/domain"
)

// Account TODO issue#131
type Account struct {
	ID        domain.ID  `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Users     []*User    `db:"-"`
}