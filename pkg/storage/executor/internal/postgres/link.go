package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/kamilsk/click/pkg/storage/types"
)

// NewLinkContext TODO issue#131
func NewLinkContext(ctx context.Context, conn *sql.Conn) linkScope {
	return linkScope{ctx, conn}
}

type linkScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO issue#131
func (scope linkScope) Create(token *types.Token, data query.CreateLink) (types.Link, error) {
	entity := types.Link{AccountID: token.User.AccountID, Name: data.Name}
	q := `INSERT INTO "link" ("id", "account_id", "name")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID, entity.Name)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a link %q",
			token.UserID, token.User.AccountID, entity.Name)
	}
	return entity, nil
}

// Read TODO issue#131
func (scope linkScope) Read(token *types.Token, data query.ReadLink) (types.Link, error) {
	entity := types.Link{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "name", "created_at", "updated_at", "deleted_at"
	        FROM "link"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the link %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// ReadByID TODO issue#131
// Deprecated: TODO issue#version3.0 use LinkEditor and gRPC gateway instead
func (scope linkScope) ReadByID(id domain.ID) (types.Link, error) {
	entity := types.Link{ID: id}
	q := `SELECT "name", "created_at", "updated_at", "deleted_at"
	        FROM "link"
	       WHERE "id" = $1`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if scanErr := row.Scan(&entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return entity, errors.NotFound(errors.LinkNotFoundMessage, scanErr, "the link %q not found", entity.ID)
		}
		return entity, errors.Database(errors.ServerErrorMessage, scanErr, "trying to populate the link %q", entity.ID)
	}
	return entity, nil
}

// ReadByAlias TODO issue#131
// Deprecated: TODO issue#logic is not transparent
func (scope linkScope) ReadByAlias(ns domain.ID, urn string) (types.Link, error) {
	var entity types.Link
	q := `SELECT "id", "name", "created_at", "updated_at", "deleted_at"
	        FROM "link"
	       WHERE "id" = (
	             SELECT "link_id"
	               FROM (
	                 SELECT "link_id"
	                   FROM "alias"
	                  WHERE "namespace_id" = $1
	                    AND "urn" = $2
	
	                  UNION
	
	                 SELECT "link_id"
	                   FROM "alias"
	                  WHERE "namespace_id" = (
	                        SELECT "account_id"
	                          FROM "namespace"
	                         WHERE "id" = $1
	                        )
	                    AND "urn" = $2
	                ) "fallback" LIMIT 1
	             )`
	row := scope.conn.QueryRowContext(scope.ctx, q, ns, urn)
	if scanErr := row.Scan(&entity.ID, &entity.Name,
		&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return entity, errors.NotFound(errors.LinkNotFoundMessage, scanErr, "the link %q not found", entity.ID)
		}
		return entity, errors.Database(errors.ServerErrorMessage, scanErr, "trying to populate the link %q", entity.ID)
	}
	return entity, nil
}

// Update TODO issue#131
func (scope linkScope) Update(token *types.Token, data query.UpdateLink) (types.Link, error) {
	entity, readErr := scope.Read(token, query.ReadLink{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	{
		entity.Name = data.Name
	}
	q := `UPDATE "link"
	         SET "name" = $1
	       WHERE "id" = $2 AND "account_id" = $3
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.Name, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the link %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO issue#131
func (scope linkScope) Delete(token *types.Token, data query.DeleteLink) (types.Link, error) {
	entity, readErr := scope.Read(token, query.ReadLink{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.Permanently {
		q := `DELETE FROM "link" WHERE "id" = $1 AND "account_id" = $2 RETURNING now()`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
		if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
			return entity, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to delete the link %q permanently",
				token.UserID, token.User.AccountID, entity.ID)
		}
		return entity, nil
	}
	q := `UPDATE "link"
	         SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the link %q safely",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
