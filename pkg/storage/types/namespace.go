package types

import (
	"time"

	"github.com/kamilsk/click/pkg/domain"
)

// Namespace TODO issue#131
type Namespace struct {
	ID        domain.ID  `db:"id"`
	AccountID domain.ID  `db:"account_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Account   *Account   `db:"-"`
}
