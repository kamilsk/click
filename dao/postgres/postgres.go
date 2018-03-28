package postgres

import (
	"database/sql"
	"encoding/json"
	"sync"

	"github.com/kamilsk/click/domain"
	"github.com/kamilsk/click/errors"
)

const dialect = "postgres"

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

	{
		var errAlias, errTarget error
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			link.Aliases, errAlias = Aliases(db, id)
		}()
		go func() {
			defer wg.Done()
			link.Targets, errTarget = Targets(db, id)
		}()
		wg.Wait()
		if errAlias != nil {
			return link, errAlias
		}
		if errTarget != nil {
			return link, errTarget
		}
	}

	return link, nil
}

// Aliases returns aliases of the Link with specified ID.
func Aliases(db *sql.DB, linkID domain.UUID) ([]domain.Alias, error) {
	aliases := make([]domain.Alias, 0, 2)
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

// Targets returns targets of the Link with specified ID.
func Targets(db *sql.DB, linkID domain.UUID) ([]domain.Target, error) {
	targets := make([]domain.Target, 0, 2)
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
			if err := json.Unmarshal(raw, &target.Rule); err != nil {
				return nil, errors.Serialization(errors.NeutralMessage, err,
					"trying to unmarshal rule of target %d of link %q", target.ID, linkID)
			}
		}
		targets = append(targets, target)
	}
	return targets, nil
}

// LinkByAlias returns the Link with its Targets and the single Alias defined by Namespace and URN.
func LinkByAlias(db *sql.DB, alias domain.Alias) (domain.Link, error) {
	var (
		link domain.Link
	)
	row := db.QueryRow(
		`SELECT "id", "link_id", "created_at", "deleted_at" FROM "alias" WHERE "namespace" = $1 AND "urn" = $2`,
		alias.Namespace, alias.URN)
	if err := row.Scan(&alias.ID, &alias.LinkID, &alias.CreatedAt, &alias.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return link, errors.NotFound(errors.LinkNotFoundMessage, err,
				"link with alias {%s:%s} not found", alias.Namespace, alias.URN)
		}
		return link, errors.Database(errors.ServerErrorMessage, err,
			"trying to populate link by alias {%s:%s}", alias.Namespace, alias.URN)
	}
	link.ID, link.Aliases = alias.LinkID, append(link.Aliases, alias)

	{
		var errLink, errTarget error
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			row := db.QueryRow(
				`SELECT "name", "status", "created_at", "updated_at" FROM "link" WHERE "id" = $1`, link.ID)
			if err := row.Scan(&link.Name, &link.Status, &link.CreatedAt, &link.UpdatedAt); err != nil {
				if err == sql.ErrNoRows {
					errLink = errors.NotFound(errors.LinkNotFoundMessage, err, "link %q not found", link.ID)
					return
				}
				errLink = errors.Database(errors.ServerErrorMessage, err, "trying to populate link %q", link.ID)
			}
		}()
		go func() {
			defer wg.Done()
			link.Targets, errTarget = Targets(db, domain.UUID(alias.LinkID))
		}()
		wg.Wait()
		if errLink != nil {
			return link, errLink
		}
		if errTarget != nil {
			return link, errTarget
		}
	}

	return link, nil
}
