package types

import (
	"time"

	"github.com/kamilsk/click/pkg/domain"
)

// Token TODO issue#131
type Token struct {
	ID        domain.ID  `db:"id"`
	UserID    domain.ID  `db:"user_id"`
	ExpiredAt *time.Time `db:"expired_at"`
	CreatedAt time.Time  `db:"created_at"`
	User      *User      `db:"-"`
}
