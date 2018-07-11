package postgres

import (
	"database/sql"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/errors"
	"golang.org/x/sync/errgroup"
)

const dialect = "postgres"

const (
	avgCount = 4
)

// Dialect returns supported database dialect.
func Dialect() string {
	return dialect
}

// Link returns the Link with its Aliases and Targets by provided ID.
func Link(db *sql.DB, id domain.UUID) (domain.Link, error) {
	var (
		link domain.Link
	)
	row := db.QueryRow(`SELECT "id", "name", "status", "created_at", "updated_at" FROM "link" WHERE "id" = $1`, id)
	if err := row.Scan(&link.ID, &link.Name, &link.Status, &link.CreatedAt, &link.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return link, errors.NotFound(errors.LinkNotFoundMessage, err, "link %q not found", id)
		}
		return link, errors.Database(errors.ServerErrorMessage, err, "trying to populate link %q", id)
	}
	g := &errgroup.Group{}
	g.Go(func() error {
		var err error
		link.Aliases, err = aliases(db, id)
		return err
	})
	g.Go(func() error {
		var err error
		link.Targets, err = targets(db, id)
		return err
	})
	if err := g.Wait(); err != nil {
		return link, err
	}
	return link, nil
}

// LinkByAlias returns the Link with its set of Alias and set of Target defined by provided namespace and URN.
func LinkByAlias(db *sql.DB, ns, urn string) (domain.Link, error) {
	var (
		aliasID   uint64
		linkID    domain.UUID
		namespace string
	)
	rows, err := db.Query(
		`SELECT "id", "link_id", "namespace" FROM "alias" WHERE "namespace" IN ($1, 'global') AND "urn" = $2`, ns, urn)
	if err != nil {
		return domain.Link{}, errors.Database(errors.ServerErrorMessage, err,
			"trying to populate link by alias {%s:%s}", ns, urn)
	}
	for rows.Next() {
		if err := rows.Scan(&aliasID, &linkID, &namespace); err != nil {
			rows.Close()
			return domain.Link{}, errors.Database(errors.ServerErrorMessage, err,
				"trying to populate link by alias {%s:%s}", ns, urn)
		}
		if namespace != "global" {
			break
		}
	}
	rows.Close()
	return Link(db, linkID)
}

// Log stores a "redirect event".
func Log(db *sql.DB, event domain.Log) (domain.Log, error) {
	encoded, err := event.Context.MarshalJSON()
	if err != nil {
		return event, errors.Serialization(errors.ServerErrorMessage, err,
			"trying to marshal a redirect event context `%#v` into JSON", event)
	}
	err = db.QueryRow(`
INSERT INTO "log" ("link_id", "alias_id", "target_id", "uri", "code", "context")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "id", "created_at"`,
		event.LinkID, event.AliasID, event.TargetID, event.URI, event.Code, encoded).Scan(&event.ID, &event.CreatedAt)
	if err != nil {
		return event, errors.Database(errors.ServerErrorMessage, err,
			"trying to insert a redirect event `%#v`", event)
	}
	return event, nil
}

// UUID returns a new generated unique identifier.
func UUID(db *sql.DB) (domain.UUID, error) {
	var id domain.UUID
	row := db.QueryRow(`SELECT uuid_generate_v4()`)
	if err := row.Scan(&id); err != nil {
		return id, errors.Database(errors.ServerErrorMessage, err, "trying to populate UUID")
	}
	return id, nil
}

func aliases(db *sql.DB, linkID domain.UUID) ([]domain.Alias, error) {
	aliases := make([]domain.Alias, 0, avgCount)
	rows, err := db.Query(
		`SELECT "id", "namespace", "urn", "created_at", "deleted_at" FROM "alias" WHERE "link_id" = $1`, linkID)
	if err != nil {
		return nil, errors.Database(errors.ServerErrorMessage, err, "trying to populate aliases of link %q", linkID)
	}
	defer rows.Close()
	for rows.Next() {
		var alias domain.Alias
		if err := rows.Scan(&alias.ID, &alias.Namespace, &alias.URN, &alias.CreatedAt, &alias.DeletedAt); err != nil {
			return nil, errors.Database(errors.ServerErrorMessage, err, "trying to populate alias of link %q", linkID)
		}
		aliases = append(aliases, alias)
	}
	return aliases, nil
}

func targets(db *sql.DB, linkID domain.UUID) ([]domain.Target, error) {
	targets := make([]domain.Target, 0, avgCount)
	rows, err := db.Query(
		`SELECT "id", "uri", "rule", "created_at", "updated_at" FROM "target" WHERE "link_id" = $1`, linkID)
	if err != nil {
		return nil, errors.Database(errors.ServerErrorMessage, err, "trying to populate targets of link %q", linkID)
	}
	defer rows.Close()
	var blob = [1024]byte{}
	for rows.Next() {
		var (
			target domain.Target
			raw    = blob[:0]
		)
		if err := rows.Scan(&target.ID, &target.URI, &raw, &target.CreatedAt, &target.UpdatedAt); err != nil {
			return nil, errors.Database(errors.ServerErrorMessage, err, "trying to populate target of link %q", linkID)
		}
		if len(raw) > 0 {
			if err := (&target.Rule).UnmarshalJSON(raw); err != nil {
				return nil, errors.Serialization(errors.NeutralMessage, err,
					"trying to unmarshal rule of target %d of link %q", target.ID, linkID)
			}
		}
		targets = append(targets, target)
	}
	return targets, nil
}
