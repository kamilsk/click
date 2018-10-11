package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/kamilsk/click/pkg/storage/types"
)

// NewLogContext TODO issue#131
func NewLogContext(ctx context.Context, conn *sql.Conn) logScope {
	return logScope{ctx, conn}
}

type logScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Write TODO issue#131
func (scope logScope) Write(data query.WriteLog) (types.Log, error) {
	entity := types.Log{
		LinkID:     data.LinkID,
		AliasID:    data.AliasID,
		TargetID:   data.TargetID,
		Identifier: data.Identifier,
		URI:        data.URI,
		Code:       data.Code,
		Context:    data.RedirectContext,
	}
	encoded, encodeErr := json.Marshal(entity.Context)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"trying to marshal context `%#v` of the redirect %q into JSON",
			entity.Context, entity.URI)
	}
	q := `INSERT INTO "log" ("account_id", "link_id", "alias_id", "target_id", "identifier", "uri", "code", "context")
	      VALUES ((SELECT "account_id" FROM "link" WHERE "id" = $1), $1, $2, $3, $4, $5, $6, $7)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q,
		entity.LinkID, entity.AliasID, entity.TargetID,
		entity.Identifier, entity.URI, entity.Code, encoded)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"trying to insert log `%s` of the redirect %q (%#+v)",
			encoded, entity.URI, data)
	}
	return entity, nil
}
