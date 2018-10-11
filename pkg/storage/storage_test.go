package storage_test

import (
	"testing"

	"github.com/kamilsk/click/pkg/errors"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/click/pkg/storage"
)

func TestMust(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { Must(func(*Storage) error { return nil }) })
	})
	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() { Must(func(*Storage) error { return errors.Simple("test") }) })
	})
}
