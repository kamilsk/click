package types

import (
	"time"

	"github.com/kamilsk/click/pkg/domain"
)

// Log TODO issue#131
type Log struct {
	ID          uint64                 `db:"id"`
	AccountID   domain.ID              `db:"account_id"`
	NamespaceID domain.ID              `db:"namespace_id"`
	LinkID      *domain.ID             `db:"link_id"`
	AliasID     *domain.ID             `db:"alias_id"`
	TargetID    *domain.ID             `db:"target_id"`
	Code        int                    `db:"code"`
	URL         string                 `db:"url"`
	Identifier  domain.ID              `db:"identifier"`
	Context     domain.RedirectContext `db:"context"`
	CreatedAt   time.Time              `db:"created_at"`
	Account     *Account               `db:"-"`
	Namespace   *Namespace             `db:"-"`
	Link        *Link                  `db:"-"`
	Alias       *Alias                 `db:"-"`
	Target      *Target                `db:"-"`
}
