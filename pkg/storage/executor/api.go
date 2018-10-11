package executor

import (
	"context"
	"database/sql"
	"fmt"

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
		exec.factory.NewAliasEditor = func(ctx context.Context, conn *sql.Conn) AliasEditor {
			return postgres.NewAliasContext(ctx, conn)
		}
		exec.factory.NewLinkEditor = func(ctx context.Context, conn *sql.Conn) LinkEditor {
			return postgres.NewLinkContext(ctx, conn)
		}
		exec.factory.NewNamespaceEditor = func(ctx context.Context, conn *sql.Conn) NamespaceEditor {
			return postgres.NewNamespaceContext(ctx, conn)
		}
		exec.factory.NewTargetEditor = func(ctx context.Context, conn *sql.Conn) TargetEditor {
			return postgres.NewTargetContext(ctx, conn)
		}
		exec.factory.NewUserManager = func(ctx context.Context, conn *sql.Conn) UserManager {
			return postgres.NewUserContext(ctx, conn)
		}
		// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
		exec.factory.NewLinkReader = func(ctx context.Context, conn *sql.Conn) LinkReader {
			return postgres.NewLinkContext(ctx, conn)
		}
	case mysqlDialect:
		fallthrough
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// AliasEditor TODO issue#131
type AliasEditor interface {
	Create(*types.Token, query.CreateAlias) (types.Alias, error)
	Read(*types.Token, query.ReadAlias) (types.Alias, error)
	Update(*types.Token, query.UpdateAlias) (types.Alias, error)
	Delete(*types.Token, query.DeleteAlias) (types.Alias, error)
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

// TargetEditor TODO issue#131
type TargetEditor interface {
	Create(*types.Token, query.CreateTarget) (types.Target, error)
	Read(*types.Token, query.ReadTarget) (types.Target, error)
	Update(*types.Token, query.UpdateTarget) (types.Target, error)
	Delete(*types.Token, query.DeleteTarget) (types.Target, error)
}

// LinkReader TODO issue#131
// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
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
		NewAliasEditor     func(context.Context, *sql.Conn) AliasEditor
		NewLinkEditor      func(context.Context, *sql.Conn) LinkEditor
		NewNamespaceEditor func(context.Context, *sql.Conn) NamespaceEditor
		NewTargetEditor    func(context.Context, *sql.Conn) TargetEditor
		NewUserManager     func(context.Context, *sql.Conn) UserManager

		// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
		NewLinkReader func(context.Context, *sql.Conn) LinkReader
	}
}

// Dialect TODO issue#131
func (e *Executor) Dialect() string {
	return e.dialect
}

// AliasEditor TODO issue#131
func (e *Executor) AliasEditor(ctx context.Context, conn *sql.Conn) AliasEditor {
	return e.factory.NewAliasEditor(ctx, conn)
}

// LinkEditor TODO issue#131
func (e *Executor) LinkEditor(ctx context.Context, conn *sql.Conn) LinkEditor {
	return e.factory.NewLinkEditor(ctx, conn)
}

// NamespaceEditor TODO issue#131
func (e *Executor) NamespaceEditor(ctx context.Context, conn *sql.Conn) NamespaceEditor {
	return e.factory.NewNamespaceEditor(ctx, conn)
}

// TargetEditor TODO issue#131
func (e *Executor) TargetEditor(ctx context.Context, conn *sql.Conn) TargetEditor {
	return e.factory.NewTargetEditor(ctx, conn)
}

// UserManager TODO issue#131
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}

// LinkReader TODO issue#131
// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
func (e *Executor) LinkReader(ctx context.Context, conn *sql.Conn) LinkReader {
	return e.factory.NewLinkReader(ctx, conn)
}
