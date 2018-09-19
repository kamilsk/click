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

		UserManager(context.Context, *sql.Conn) executor.UserManager
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
