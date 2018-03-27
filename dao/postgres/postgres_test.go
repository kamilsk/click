//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package postgres_test -destination $PWD/dao/postgres/mock_db_test.go database/sql/driver Conn,Driver,Stmt,Rows
package postgres_test

import (
	"testing"

	"github.com/kamilsk/click/dao/postgres"
	"github.com/kamilsk/click/domain"
	"github.com/stretchr/testify/assert"
)

const (
	DSN  = "postgres://postgres:postgres@db:5432/postgres"
	UUID = domain.UUID("41ca5e09-3ce2-4094-b108-3ecc257c6fa4")
)

func TestDialect(t *testing.T) {
	assert.Equal(t, "postgres", postgres.Dialect())
}
