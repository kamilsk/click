package postgres_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/storage/executor"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	. "github.com/kamilsk/click/pkg/storage/executor/internal/postgres"
)

func TestLogWriter(t *testing.T) {
	id := domain.ID("10000000-2000-4000-8000-160000000000")
	t.Run("write", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer conn.Close()

			mock.
				ExpectQuery(`INSERT INTO "event"`).
				WithArgs(id, &id, &id, &id, http.StatusFound, "test", id, []byte(`{}`)).
				WillReturnRows(
					sqlmock.
						NewRows([]string{"id", "created_at"}).
						AddRow(uint64(1), time.Now()),
				)

			var exec executor.LogWriter = NewLogContext(ctx, conn)
			log, err := exec.Write(query.WriteLog{
				RedirectEvent: domain.RedirectEvent{
					NamespaceID: id, LinkID: &id, AliasID: &id, TargetID: &id,
					Code: http.StatusFound, URL: "test", Identifier: id,
					Context: domain.RedirectContext{},
				},
			})
			assert.NoError(t, err)
			assert.Equal(t, uint64(1), log.ID)
			assert.NotEmpty(t, log.CreatedAt)
		})
		t.Run("serialization error", func(t *testing.T) {
			// TODO issue#126
		})
		t.Run("database error", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer conn.Close()

			mock.
				ExpectQuery(`INSERT INTO "event"`).
				WithArgs(id, &id, &id, &id, http.StatusFound, "test", id, []byte(`{}`)).
				WillReturnError(errors.Simple("test"))

			var exec executor.LogWriter = NewLogContext(ctx, conn)
			log, err := exec.Write(query.WriteLog{
				RedirectEvent: domain.RedirectEvent{
					NamespaceID: id, LinkID: &id, AliasID: &id, TargetID: &id,
					Code: http.StatusFound, URL: "test", Identifier: id,
					Context: domain.RedirectContext{},
				},
			})
			assert.Error(t, err)
			assert.Empty(t, log.ID)
			assert.Empty(t, log.CreatedAt)
		})
	})
}
