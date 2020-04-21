package postgres_test

import (
	"context"
	"testing"

	"go.octolab.org/ecosystem/click/internal/storage/executor"
	. "go.octolab.org/ecosystem/click/internal/storage/executor/internal/postgres"
)

func TestTargetEditor(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TargetEditor = NewTargetContext(ctx, nil)
	})
	t.Run("read", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TargetEditor = NewTargetContext(ctx, nil)
	})
	t.Run("update", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TargetEditor = NewTargetContext(ctx, nil)
	})
	t.Run("delete", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TargetEditor = NewTargetContext(ctx, nil)
	})
}
