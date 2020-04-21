package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"go.octolab.org/ecosystem/click/internal/errors"
	"go.octolab.org/ecosystem/click/internal/storage/query"
	"go.octolab.org/ecosystem/click/internal/storage/types"
)

// NewEventContext TODO issue#131
func NewEventContext(ctx context.Context, conn *sql.Conn) eventScope {
	return eventScope{ctx, conn}
}

type eventScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Write TODO issue#131
func (scope eventScope) Write(data query.WriteLog) (types.Event, error) {
	entity := types.Event{
		NamespaceID: data.NamespaceID,
		LinkID:      data.LinkID,
		AliasID:     data.AliasID,
		TargetID:    data.TargetID,
		Identifier:  data.Identifier,
		Context:     data.Context,
		Code:        data.Code,
		URL:         data.URL,
	}
	encoded, encodeErr := json.Marshal(entity.Context)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"trying to marshal context `%#v` of the redirect %q into JSON",
			entity.Context, entity.URL)
	}
	q := `INSERT INTO "event"
	      ("account_id", "namespace_id", "link_id", "alias_id", "target_id", "identifier", "context", "code", "url")
	      VALUES ((SELECT "account_id" FROM "namespace" WHERE "id" = $1), $1, $2, $3, $4, $5, $6, $7, $8)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q,
		entity.NamespaceID, entity.LinkID, entity.AliasID, entity.TargetID, entity.Identifier,
		encoded, entity.Code, entity.URL,
	)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"trying to insert event `%s` of the redirect %q (%+v)",
			encoded, entity.URL, data)
	}
	return entity, nil
}
