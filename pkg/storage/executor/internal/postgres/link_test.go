package postgres_test

import (
	"context"
	"testing"

	"github.com/kamilsk/click/pkg/storage/executor"
	. "github.com/kamilsk/click/pkg/storage/executor/internal/postgres"
)

func TestLinkEditor(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.LinkEditor = NewLinkContext(ctx, nil)
	})
	t.Run("read", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.LinkEditor = NewLinkContext(ctx, nil)
	})
	t.Run("update", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.LinkEditor = NewLinkContext(ctx, nil)
	})
	t.Run("delete", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.LinkEditor = NewLinkContext(ctx, nil)
	})
}

func TestLinkReader(t *testing.T) {
	t.Run("read by ID", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.LinkReader = NewLinkContext(ctx, nil)
	})
}
