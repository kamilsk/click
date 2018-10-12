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
		NamespaceID: data.NamespaceID,
		LinkID:      data.LinkID,
		AliasID:     data.AliasID,
		TargetID:    data.TargetID,
		Code:        data.Code,
		URL:         data.URL,
		Identifier:  data.Identifier,
		Context:     data.Context,
	}
	encoded, encodeErr := json.Marshal(entity.Context)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"trying to marshal context `%#v` of the redirect %q into JSON",
			entity.Context, entity.URL)
	}
	q := `INSERT INTO "event" ("account_id", "namespace_id", "link_id", "alias_id", "target_id", "code", "url", "identifier", "context")
	      VALUES ((SELECT "account_id" FROM "namespace" WHERE "id" = $1), $1, $2, $3, $4, $5, $6, $7, $8)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q,
		entity.NamespaceID, entity.LinkID, entity.AliasID, entity.TargetID,
		entity.Code, entity.URL, entity.Identifier, encoded)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"trying to insert log `%s` of the redirect %q (%#+v)",
			encoded, entity.URL, data)
	}
	return entity, nil
}
