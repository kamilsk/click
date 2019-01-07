package executor_test

import (
	"context"
	"database/sql"
	"testing"

	. "github.com/kamilsk/click/pkg/storage/executor"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type contract interface {
		Dialect() string

		AliasEditor(context.Context, *sql.Conn) AliasEditor
		AliasReader(context.Context, *sql.Conn) AliasReader
		LinkEditor(context.Context, *sql.Conn) LinkEditor
		LinkReader(context.Context, *sql.Conn) LinkReader
		LogWriter(context.Context, *sql.Conn) LogWriter
		NamespaceEditor(context.Context, *sql.Conn) NamespaceEditor
		TargetEditor(context.Context, *sql.Conn) TargetEditor
		TargetReader(context.Context, *sql.Conn) TargetReader
		UserManager(context.Context, *sql.Conn) UserManager
	}
	t.Run("PostgreSQL", func(t *testing.T) {
		assert.NotPanics(t, func() {
			dialect, ctx := "postgres", context.Background()
			var exec contract = New(dialect)
			assert.Equal(t, dialect, exec.Dialect())

			assert.NotNil(t, exec.AliasEditor(ctx, nil))
			assert.NotNil(t, exec.AliasReader(ctx, nil))
			assert.NotNil(t, exec.LinkEditor(ctx, nil))
			assert.NotNil(t, exec.LinkReader(ctx, nil))
			assert.NotNil(t, exec.LogWriter(ctx, nil))
			assert.NotNil(t, exec.NamespaceEditor(ctx, nil))
			assert.NotNil(t, exec.TargetEditor(ctx, nil))
			assert.NotNil(t, exec.TargetReader(ctx, nil))
			assert.NotNil(t, exec.UserManager(ctx, nil))

			assert.NotNil(t, exec.LinkReader(ctx, nil))
		})
	})
	t.Run("MySQL", func(t *testing.T) {
		assert.Panics(t, func() {
			dialect := "mysql"
			var _ contract = New(dialect)
		})
	})
}
