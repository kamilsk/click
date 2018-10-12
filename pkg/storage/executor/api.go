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
		exec.factory.NewAliasReader = func(ctx context.Context, conn *sql.Conn) AliasReader {
			return postgres.NewAliasContext(ctx, conn)
		}
		exec.factory.NewLinkEditor = func(ctx context.Context, conn *sql.Conn) LinkEditor {
			return postgres.NewLinkContext(ctx, conn)
		}
		exec.factory.NewLinkReader = func(ctx context.Context, conn *sql.Conn) LinkReader {
			return postgres.NewLinkContext(ctx, conn)
		}
		exec.factory.NewLogWriter = func(ctx context.Context, conn *sql.Conn) LogWriter {
			return postgres.NewLogContext(ctx, conn)
		}
		exec.factory.NewNamespaceEditor = func(ctx context.Context, conn *sql.Conn) NamespaceEditor {
			return postgres.NewNamespaceContext(ctx, conn)
		}
		exec.factory.NewTargetEditor = func(ctx context.Context, conn *sql.Conn) TargetEditor {
			return postgres.NewTargetContext(ctx, conn)
		}
		exec.factory.NewTargetReader = func(ctx context.Context, conn *sql.Conn) TargetReader {
			return postgres.NewTargetContext(ctx, conn)
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

// AliasEditor TODO issue#131
type AliasEditor interface {
	// Create TODO issue#131
	Create(*types.Token, query.CreateAlias) (types.Alias, error)
	// Read TODO issue#131
	Read(*types.Token, query.ReadAlias) (types.Alias, error)
	// Update TODO issue#131
	Update(*types.Token, query.UpdateAlias) (types.Alias, error)
	// Delete TODO issue#131
	Delete(*types.Token, query.DeleteAlias) (types.Alias, error)
}

// AliasReader TODO issue#131
type AliasReader interface {
	// ReadAllByLink TODO issue#131
	ReadAllByLink(domain.ID) ([]types.Alias, error)
}

// LinkEditor TODO issue#131
type LinkEditor interface {
	// Create TODO issue#131
	Create(*types.Token, query.CreateLink) (types.Link, error)
	// Read TODO issue#131
	Read(*types.Token, query.ReadLink) (types.Link, error)
	// Update TODO issue#131
	Update(*types.Token, query.UpdateLink) (types.Link, error)
	// Delete TODO issue#131
	Delete(*types.Token, query.DeleteLink) (types.Link, error)
}

// LinkReader TODO issue#131
type LinkReader interface {
	// ReadByID TODO issue#131
	// Deprecated: TODO issue#version3.0 use LinkEditor and gRPC gateway instead
	ReadByID(domain.ID) (types.Link, error)
	// ReadByAlias TODO issue#131
	// Deprecated: TODO issue#refactoring logic is not transparent
	ReadByAlias(ns domain.ID, urn string) (types.Link, error)
}

// LogWriter TODO issue#131
type LogWriter interface {
	// Write TODO issue#131
	Write(query.WriteLog) (types.Log, error)
}

// NamespaceEditor TODO issue#131
type NamespaceEditor interface {
	// Create TODO issue#131
	Create(*types.Token, query.CreateNamespace) (types.Namespace, error)
	// Read TODO issue#131
	Read(*types.Token, query.ReadNamespace) (types.Namespace, error)
	// Update TODO issue#131
	Update(*types.Token, query.UpdateNamespace) (types.Namespace, error)
	// Delete TODO issue#131
	Delete(*types.Token, query.DeleteNamespace) (types.Namespace, error)
}

// TargetEditor TODO issue#131
type TargetEditor interface {
	// Create TODO issue#131
	Create(*types.Token, query.CreateTarget) (types.Target, error)
	// Read TODO issue#131
	Read(*types.Token, query.ReadTarget) (types.Target, error)
	// Update TODO issue#131
	Update(*types.Token, query.UpdateTarget) (types.Target, error)
	// Delete TODO issue#131
	Delete(*types.Token, query.DeleteTarget) (types.Target, error)
}

// TargetReader TODO issue#131
type TargetReader interface {
	// ReadAllByLink TODO issue#131
	ReadAllByLink(link domain.ID) ([]types.Target, error)
}

// UserManager TODO issue#131
type UserManager interface {
	// Token TODO issue#131
	Token(domain.ID) (*types.Token, error)
}

// Executor TODO issue#131
type Executor struct {
	dialect string
	factory struct {
		NewAliasEditor     func(context.Context, *sql.Conn) AliasEditor
		NewAliasReader     func(context.Context, *sql.Conn) AliasReader
		NewLinkEditor      func(context.Context, *sql.Conn) LinkEditor
		NewLinkReader      func(context.Context, *sql.Conn) LinkReader
		NewLogWriter       func(context.Context, *sql.Conn) LogWriter
		NewNamespaceEditor func(context.Context, *sql.Conn) NamespaceEditor
		NewTargetEditor    func(context.Context, *sql.Conn) TargetEditor
		NewTargetReader    func(context.Context, *sql.Conn) TargetReader
		NewUserManager     func(context.Context, *sql.Conn) UserManager
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

// AliasReader TODO issue#131
func (e *Executor) AliasReader(ctx context.Context, conn *sql.Conn) AliasReader {
	return e.factory.NewAliasReader(ctx, conn)
}

// LinkEditor TODO issue#131
func (e *Executor) LinkEditor(ctx context.Context, conn *sql.Conn) LinkEditor {
	return e.factory.NewLinkEditor(ctx, conn)
}

// LinkReader TODO issue#131
func (e *Executor) LinkReader(ctx context.Context, conn *sql.Conn) LinkReader {
	return e.factory.NewLinkReader(ctx, conn)
}

// LogWriter TODO issue#131
func (e *Executor) LogWriter(ctx context.Context, conn *sql.Conn) LogWriter {
	return e.factory.NewLogWriter(ctx, conn)
}

// NamespaceEditor TODO issue#131
func (e *Executor) NamespaceEditor(ctx context.Context, conn *sql.Conn) NamespaceEditor {
	return e.factory.NewNamespaceEditor(ctx, conn)
}

// TargetEditor TODO issue#131
func (e *Executor) TargetEditor(ctx context.Context, conn *sql.Conn) TargetEditor {
	return e.factory.NewTargetEditor(ctx, conn)
}

// TargetReader TODO issue#131
func (e *Executor) TargetReader(ctx context.Context, conn *sql.Conn) TargetReader {
	return e.factory.NewTargetReader(ctx, conn)
}

// UserManager TODO issue#131
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}
