package postgres

import (
	"database/sql"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/errors"
)

// Log stores a "redirect event".
func Log(db *sql.DB, event domain.Redirect) (domain.Redirect, error) {
	encoded, err := event.Context.MarshalJSON()
	if err != nil {
		return event, errors.Serialization(errors.ServerErrorMessage, err,
			"trying to marshal a redirect event context `%#v` into JSON", event)
	}
	err = db.QueryRow(`
INSERT INTO "log" ("link_id", "alias_id", "target_id", "uri", "code", "context")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "id"`,
		event.LinkID, event.AliasID, event.TargetID, event.URI, event.Code, encoded).Scan(&event.ID)
	if err != nil {
		return event, errors.Database(errors.ServerErrorMessage, err,
			"trying to insert a redirect event `%#v`", event)
	}
	return event, nil
}
