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
		UserManager(context.Context, *sql.Conn) executor.UserManager

		// Deprecated TODO issue#version3.0 use LinkEditor and gRPC gateway instead
		LinkReader(context.Context, *sql.Conn) executor.LinkReader
	}
	t.Run("PostgreSQL", func(t *testing.T) {
		assert.NotPanics(t, func() {
			var exec contract = executor.New("postgres")
			assert.NotEmpty(t, exec.Dialect())
			assert.NotNil(t, exec.UserManager(nil, nil))
		})
	})
	t.Run("MySQL", func(t *testing.T) {
		assert.Panics(t, func() {
			var _ contract = executor.New("mysql")
		})
	})
}
