package types

import (
	"time"

	"go.octolab.org/ecosystem/click/internal/domain"
)

// Target TODO issue#131
type Target struct {
	ID         domain.ID         `db:"id"`
	AccountID  domain.ID         `db:"account_id"`
	LinkID     domain.ID         `db:"link_id"`
	URL        string            `db:"url"`
	Rule       domain.Rule       `db:"rule,deprecated"`
	BinaryRule domain.BinaryRule `db:"b_rule"`
	CreatedAt  time.Time         `db:"created_at"`
	UpdatedAt  *time.Time        `db:"updated_at"`
	DeletedAt  *time.Time        `db:"deleted_at"`
	Account    *Account          `db:"-"`
	Link       *Link             `db:"-"`
}
