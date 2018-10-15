package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/kamilsk/click/pkg/storage/types"
)

// NewNamespaceContext TODO issue#131
func NewNamespaceContext(ctx context.Context, conn *sql.Conn) namespaceScope {
	return namespaceScope{ctx, conn}
}

type namespaceScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO issue#131
func (scope namespaceScope) Create(token *types.Token, data query.CreateNamespace) (types.Namespace, error) {
	entity := types.Namespace{AccountID: token.User.AccountID, Name: data.Name}
	q := `INSERT INTO "namespace" ("id", "account_id", "name")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID, entity.Name)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a namespace %q",
			token.UserID, token.User.AccountID, entity.Name)
	}
	return entity, nil
}

// Read TODO issue#131
func (scope namespaceScope) Read(token *types.Token, data query.ReadNamespace) (types.Namespace, error) {
	entity := types.Namespace{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "name", "created_at", "updated_at", "deleted_at"
	        FROM "namespace"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Name, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the namespace %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Update TODO issue#131
func (scope namespaceScope) Update(token *types.Token, data query.UpdateNamespace) (types.Namespace, error) {
	entity, readErr := scope.Read(token, query.ReadNamespace{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	{
		entity.Name = data.Name
	}
	q := `UPDATE "namespace"
	         SET "name" = $1
	       WHERE "id" = $2 AND "account_id" = $3
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.Name, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the namespace %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO issue#131
func (scope namespaceScope) Delete(token *types.Token, data query.DeleteNamespace) (types.Namespace, error) {
	entity, readErr := scope.Read(token, query.ReadNamespace{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.Permanently {
		q := `DELETE FROM "namespace" WHERE "id" = $1 AND "account_id" = $2 RETURNING now()`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
		if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
			return entity, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to delete the namespace %q permanently",
				token.UserID, token.User.AccountID, entity.ID)
		}
		return entity, nil
	}
	q := `UPDATE "namespace"
	         SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the namespace %q safely",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
