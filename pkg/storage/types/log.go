package types

import (
	"time"

	"github.com/kamilsk/click/pkg/domain"
)

// Log TODO issue#131
type Log struct {
	ID         uint64                 `db:"id"`
	AccountID  domain.ID              `db:"account_id"`
	LinkID     domain.ID              `db:"link_id"`
	AliasID    domain.ID              `db:"alias_id"`
	TargetID   domain.ID              `db:"target_id"`
	Identifier domain.ID              `db:"identifier"`
	URI        string                 `db:"uri"`
	Code       uint16                 `db:"code"`
	Context    domain.RedirectContext `db:"context"`
	CreatedAt  time.Time              `db:"created_at"`
	Account    *Account               `db:"-"`
	Link       *Link                  `db:"-"`
	Alias      *Alias                 `db:"-"`
	Target     *Target                `db:"-"`
}
