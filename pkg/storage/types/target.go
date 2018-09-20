package types

import (
	"time"

	"github.com/kamilsk/click/pkg/domain"
)

// Target TODO issue#131
type Target struct {
	ID         domain.ID  `db:"id"`
	AccountID  domain.ID  `db:"account_id"`
	LinkID     domain.ID  `db:"link_id"`
	URI        string     `db:"uri"`
	Rule       []byte     `db:"rule,deprecated"`
	BinaryRule []byte     `db:"b_rule"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
	Link       *Link      `db:"-"`
}
