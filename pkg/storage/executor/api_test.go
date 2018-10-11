package executor_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	
	. "github.com/kamilsk/click/pkg/storage/executor"
)

func TestNew(t *testing.T) {
	type contract interface {
		Dialect() string

		AliasEditor(context.Context, *sql.Conn) AliasEditor
		LinkEditor(context.Context, *sql.Conn) LinkEditor
		NamespaceEditor(context.Context, *sql.Conn) NamespaceEditor
		TargetEditor(context.Context, *sql.Conn) TargetEditor
		UserManager(context.Context, *sql.Conn) UserManager

		// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
		LinkReader(context.Context, *sql.Conn) LinkReader
	}
	t.Run("PostgreSQL", func(t *testing.T) {
		assert.NotPanics(t, func() {
			dialect, ctx := "postgres", context.Background()
			var exec contract = New(dialect)
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
			var _ contract = New(dialect)
		})
	})
}
