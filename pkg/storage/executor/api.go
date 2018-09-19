package executor

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/storage/executor/internal/postgres"
	"github.com/kamilsk/click/pkg/storage/types"
)

const (
	postgresDialect = "postgres"
	mysqlDialect    = "mysql"
)

// New TODO issue#131
func New(dialect string) *Executor {
	exec := &Executor{dialect: dialect}
	switch exec.dialect {
	case postgresDialect:
		exec.factory.NewUserManager = func(ctx context.Context, conn *sql.Conn) UserManager {
			return postgres.NewUserContext(ctx, conn)
		}
	case mysqlDialect:
		fallthrough
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// UserManager TODO issue#131
type UserManager interface {
	Token(domain.ID) (*types.Token, error)
}

// Executor TODO issue#131
type Executor struct {
	dialect string
	factory struct {
		NewUserManager func(context.Context, *sql.Conn) UserManager
	}
}

// Dialect TODO issue#131
func (e *Executor) Dialect() string {
	return e.dialect
}

// UserManager TODO issue#131
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}
