package executor

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/storage/executor/internal/postgres"
	"github.com/kamilsk/click/pkg/storage/query"
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
		exec.factory.NewLinkEditor = func(ctx context.Context, conn *sql.Conn) LinkEditor {
			return postgres.NewLinkContext(ctx, conn)
		}
		exec.factory.NewLinkReader = func(ctx context.Context, conn *sql.Conn) LinkReader {
			return postgres.NewLinkContext(ctx, conn)
		}
		exec.factory.NewNamespaceEditor = func(ctx context.Context, conn *sql.Conn) NamespaceEditor {
			return postgres.NewNamespaceContext(ctx, conn)
		}
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

// LinkEditor TODO issue#131
type LinkEditor interface {
	Create(*types.Token, query.CreateLink) (types.Link, error)
	Read(*types.Token, query.ReadLink) (types.Link, error)
	Update(*types.Token, query.UpdateLink) (types.Link, error)
	Delete(*types.Token, query.DeleteLink) (types.Link, error)
}

// NamespaceEditor TODO issue#131
type NamespaceEditor interface {
	Create(*types.Token, query.CreateNamespace) (types.Namespace, error)
	Read(*types.Token, query.ReadNamespace) (types.Namespace, error)
	Update(*types.Token, query.UpdateNamespace) (types.Namespace, error)
	Delete(*types.Token, query.DeleteNamespace) (types.Namespace, error)
}

// LinkReader TODO issue#131
type LinkReader interface {
	ReadByID(domain.ID) (types.Link, error)
}

// UserManager TODO issue#131
type UserManager interface {
	Token(domain.ID) (*types.Token, error)
}

// Executor TODO issue#131
type Executor struct {
	dialect string
	factory struct {
		NewLinkEditor      func(context.Context, *sql.Conn) LinkEditor
		NewLinkReader      func(context.Context, *sql.Conn) LinkReader
		NewNamespaceEditor func(context.Context, *sql.Conn) NamespaceEditor
		NewUserManager     func(context.Context, *sql.Conn) UserManager
	}
}

// Dialect TODO issue#131
func (e *Executor) Dialect() string {
	return e.dialect
}

// LinkEditor TODO issue#131
func (e *Executor) LinkEditor(ctx context.Context, conn *sql.Conn) LinkEditor {
	return e.factory.NewLinkEditor(ctx, conn)
}

// NamespaceEditor TODO issue#131
func (e *Executor) NamespaceEditor(ctx context.Context, conn *sql.Conn) NamespaceEditor {
	return e.factory.NewNamespaceEditor(ctx, conn)
}

// LinkReader TODO issue#131
func (e *Executor) LinkReader(ctx context.Context, conn *sql.Conn) LinkReader {
	return e.factory.NewLinkReader(ctx, conn)
}

// UserManager TODO issue#131
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}
