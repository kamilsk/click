package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/kamilsk/click/pkg/storage/types"
)

// NewAliasContext TODO issue#131
func NewAliasContext(ctx context.Context, conn *sql.Conn) aliasScope {
	return aliasScope{ctx, conn}
}

type aliasScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO issue#131
func (scope aliasScope) Create(token *types.Token, data query.CreateAlias) (types.Alias, error) {
	entity := types.Alias{
		AccountID:   token.User.AccountID,
		LinkID:      data.LinkID,
		NamespaceID: data.NamespaceID,
		URN:         data.URN,
	}
	q := `INSERT INTO "alias" ("id", "account_id", "link_id", "namespace_id", "urn") VALUES ($1, $2, $3, $4, $5)
	      RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID,
		entity.LinkID, entity.NamespaceID, entity.URN)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create an alias %q",
			token.UserID, token.User.AccountID, entity.URN)
	}
	return entity, nil
}

// Read TODO issue#131
func (scope aliasScope) Read(token *types.Token, data query.ReadAlias) (types.Alias, error) {
	entity := types.Alias{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "link_id", "namespace_id", "urn", "created_at", "updated_at", "deleted_at" FROM "alias"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.LinkID, &entity.NamespaceID, &entity.URN,
		&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the alias %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Update TODO issue#131
func (scope aliasScope) Update(token *types.Token, data query.UpdateAlias) (types.Alias, error) {
	entity, readErr := scope.Read(token, query.ReadAlias{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.URN != "" {
		entity.URN = data.URN
	}
	q := `UPDATE "alias" SET "urn" = $1
	       WHERE "id" = $2 AND "account_id" = $3
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.URN, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the alias %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO issue#131
func (scope aliasScope) Delete(token *types.Token, data query.DeleteAlias) (types.Alias, error) {
	entity, readErr := scope.Read(token, query.ReadAlias{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.Permanently {
		q := `DELETE FROM "alias" WHERE "id" = $1 AND "account_id" = $2 RETURNING now()`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
		if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
			return entity, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to delete the alias %q permanently",
				token.UserID, token.User.AccountID, entity.ID)
		}
		return entity, nil
	}
	q := `UPDATE "alias" SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the alias %q safely",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
