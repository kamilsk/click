package types

import (
	"time"

	"go.octolab.org/ecosystem/click/internal/domain"
)

// Alias TODO issue#131
type Alias struct {
	ID          domain.ID  `db:"id"`
	AccountID   domain.ID  `db:"account_id"`
	LinkID      domain.ID  `db:"link_id"`
	NamespaceID domain.ID  `db:"namespace_id"`
	URN         string     `db:"urn"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	Account     *Account   `db:"-"`
	Link        *Link      `db:"-"`
	Namespace   *Namespace `db:"-"`
}
