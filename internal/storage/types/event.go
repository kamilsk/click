package types

import (
	"time"

	"go.octolab.org/ecosystem/click/internal/domain"
)

// Event TODO issue#131
type Event struct {
	ID          uint64                 `db:"id"`
	AccountID   domain.ID              `db:"account_id"`
	NamespaceID domain.ID              `db:"namespace_id"`
	LinkID      *domain.ID             `db:"link_id"`
	AliasID     *domain.ID             `db:"alias_id"`
	TargetID    *domain.ID             `db:"target_id"`
	Identifier  *domain.ID             `db:"identifier"`
	Context     domain.RedirectContext `db:"context"`
	Code        int                    `db:"code"`
	URL         string                 `db:"url"`
	CreatedAt   time.Time              `db:"created_at"`
	Account     *Account               `db:"-"`
	Namespace   *Namespace             `db:"-"`
	Link        *Link                  `db:"-"`
	Alias       *Alias                 `db:"-"`
	Target      *Target                `db:"-"`
}
