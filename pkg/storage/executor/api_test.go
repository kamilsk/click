package executor_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/kamilsk/click/pkg/storage/executor"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type contract interface {
		Dialect() string

		AliasEditor(context.Context, *sql.Conn) executor.AliasEditor
		LinkEditor(context.Context, *sql.Conn) executor.LinkEditor
		NamespaceEditor(context.Context, *sql.Conn) executor.NamespaceEditor
		TargetEditor(context.Context, *sql.Conn) executor.TargetEditor
		UserManager(context.Context, *sql.Conn) executor.UserManager

		// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
		LinkReader(context.Context, *sql.Conn) executor.LinkReader
	}
	t.Run("PostgreSQL", func(t *testing.T) {
		assert.NotPanics(t, func() {
			dialect, ctx := "postgres", context.Background()
			var exec contract = executor.New(dialect)
			assert.Equal(t, dialect, exec.Dialect())

			assert.NotNil(t, exec.AliasEditor(ctx, nil))
			assert.NotNil(t, exec.LinkEditor(ctx, nil))
			assert.NotNil(t, exec.NamespaceEditor(ctx, nil))
			assert.NotNil(t, exec.TargetEditor(ctx, nil))
			assert.NotNil(t, exec.UserManager(ctx, nil))

			assert.NotNil(t, exec.LinkReader(ctx, nil))
		})
	})
	t.Run("MySQL", func(t *testing.T) {
		assert.Panics(t, func() {
			dialect := "mysql"
			var _ contract = executor.New(dialect)
		})
	})
}
