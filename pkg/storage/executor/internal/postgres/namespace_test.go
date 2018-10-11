package postgres_test

import (
	"context"
	"testing"

	"github.com/kamilsk/click/pkg/storage/executor"

	. "github.com/kamilsk/click/pkg/storage/executor/internal/postgres"
)

func TestNamespaceEditor(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.NamespaceEditor = NewNamespaceContext(ctx, nil)
	})
	t.Run("read", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.NamespaceEditor = NewNamespaceContext(ctx, nil)
	})
	t.Run("update", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.NamespaceEditor = NewNamespaceContext(ctx, nil)
	})
	t.Run("delete", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.NamespaceEditor = NewNamespaceContext(ctx, nil)
	})
}
