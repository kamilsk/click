package storage

import "time"

// Account TODO
type Account struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Users     []*User    `db:"-"`
}

// User TODO
type User struct {
	ID        string     `db:"id"`
	AccountID string     `db:"account_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Account   *Account   `db:"-"`
	Tokens    []*Token   `db:"-"`
}

// Token TODO
type Token struct {
	ID        string     `db:"id"`
	UserID    string     `db:"user_id"`
	ExpiredAt *time.Time `db:"expired_at"`
	CreatedAt time.Time  `db:"created_at"`
	User      *User      `db:"-"`
}

// Namespace TODO
type Namespace struct {
	ID        string     `db:"id"`
	AccountID string     `db:"account_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Account   *Account   `db:"-"`
}

// Link TODO
type Link struct {
	ID        string     `db:"id"`
	AccountID string     `db:"account_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Account   *Account   `db:"-"`
}

// Alias TODO
type Alias struct {
	ID          string     `db:"id"`
	LinkID      string     `db:"link_id"`
	NamespaceID string     `db:"namespace_id"`
	URN         string     `db:"urn"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	Link        *Link      `db:"-"`
	Namespace   *Namespace `db:"-"`
}

// Target TODO
type Target struct {
	ID         string     `db:"id"`
	LinkID     string     `db:"link_id"`
	URI        string     `db:"uri"`
	Rule       []byte     `db:"rule,deprecated"`
	BinaryRule []byte     `db:"b_rule"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
	Link       *Link      `db:"-"`
}

// Log TODO
type Log struct {
	ID         uint64    `db:"id"`
	AccountID  string    `db:"account_id"`
	LinkID     string    `db:"link_id"`
	AliasID    string    `db:"alias_id"`
	TargetID   string    `db:"target_id"`
	Identifier string    `db:"identifier"`
	URI        string    `db:"uri"`
	Code       uint16    `db:"code"`
	Context    []byte    `db:"context"`
	CreatedAt  time.Time `db:"created_at"`
	Account    *Account  `db:"-"`
	Link       *Link     `db:"-"`
	Alias      *Alias    `db:"-"`
	Target     *Target   `db:"-"`
}
