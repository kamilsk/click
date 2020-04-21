package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"go.octolab.org/ecosystem/click/internal/domain"
	"go.octolab.org/ecosystem/click/internal/errors"
	"go.octolab.org/ecosystem/click/internal/storage/query"
	"go.octolab.org/ecosystem/click/internal/storage/types"
)

// NewTargetContext TODO issue#131
func NewTargetContext(ctx context.Context, conn *sql.Conn) targetScope {
	return targetScope{ctx, conn}
}

type targetScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO issue#131
func (scope targetScope) Create(token *types.Token, data query.CreateTarget) (types.Target, error) {
	entity := types.Target{
		AccountID:  token.User.AccountID,
		LinkID:     data.LinkID,
		URL:        data.URL,
		Rule:       data.Rule,
		BinaryRule: data.BinaryRule,
	}
	encodedRule, encodeErr := json.Marshal(entity.Rule)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"user %q of account %q tried to marshal a target rule `%#v` into JSON",
			token.UserID, token.User.AccountID, entity.Rule)
	}
	encodedBinaryRule := []byte(entity.BinaryRule)
	q := `INSERT INTO "target" ("id", "account_id", "link_id", "url", "rule", "b_rule")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3, $4, $5, $6)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID, entity.LinkID,
		entity.URL, encodedRule, encodedBinaryRule)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a target %q",
			token.UserID, token.User.AccountID, entity.URL)
	}
	return entity, nil
}

// Read TODO issue#131
func (scope targetScope) Read(token *types.Token, data query.ReadTarget) (types.Target, error) {
	var encodedRule, encodedBinaryRule []byte
	entity := types.Target{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "link_id", "url", "rule", "b_rule", "created_at", "updated_at", "deleted_at"
	        FROM "target"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.LinkID, &entity.URL,
		&encodedRule, &encodedBinaryRule,
		&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the target %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	if err := json.Unmarshal(encodedRule, &entity.Rule); err != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, err,
			"trying to unmarshal JSON `%s` of the target rule %q",
			encodedRule, entity.ID)
	}
	entity.BinaryRule = domain.BinaryRule(encodedBinaryRule)
	return entity, nil
}

// ReadAllByLink TODO issue#131
func (scope targetScope) ReadAllByLink(link domain.ID) ([]types.Target, error) {
	q := `SELECT "id", "url", "rule", "b_rule", "created_at", "updated_at", "deleted_at"
	        FROM "target"
	       WHERE "link_id" = $1 AND "deleted_at" IS NULL`
	rows, queryErr := scope.conn.QueryContext(scope.ctx, q, link)
	if queryErr != nil {
		return nil, errors.Database(errors.ServerErrorMessage, queryErr,
			"trying to read targets of the link %q", link)
	}
	defer func() { _ = rows.Close() }()
	result := make([]types.Target, 0, 4)
	for rows.Next() {
		var encodedRule, encodedBinaryRule []byte
		entity := types.Target{LinkID: link}
		if scanErr := rows.Scan(&entity.ID, &entity.URL,
			&encodedRule, &encodedBinaryRule,
			&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); scanErr != nil {
			return nil, errors.Database(errors.ServerErrorMessage, scanErr,
				"trying to read a target of the link %q", link)
		}
		if err := json.Unmarshal(encodedRule, &entity.Rule); err != nil {
			return nil, errors.Serialization(errors.ServerErrorMessage, err,
				"trying to unmarshal JSON `%s` of the target rule %q",
				encodedRule, entity.ID)
		}
		entity.BinaryRule = domain.BinaryRule(encodedBinaryRule)
		result = append(result, entity)
	}
	if loopErr := rows.Err(); loopErr != nil {
		return nil, errors.Database(errors.ServerErrorMessage, loopErr,
			"trying to read targets of the link %q", link)
	}
	return result, nil
}

// Update TODO issue#131
func (scope targetScope) Update(token *types.Token, data query.UpdateTarget) (types.Target, error) {
	entity, readErr := scope.Read(token, query.ReadTarget{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	{
		entity.LinkID = data.LinkID
		entity.URL = data.URL
		entity.Rule = data.Rule
		entity.BinaryRule = data.BinaryRule
	}
	encodedRule, encodeErr := json.Marshal(entity.Rule)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"user %q of account %q tried to marshal rule `%#v` of the target %q into JSON",
			token.UserID, token.User.AccountID, entity.Rule, entity.ID)
	}
	encodedBinaryRule := []byte(entity.BinaryRule)
	q := `UPDATE "target"
	         SET "link_id" = $1, "url" = $2, "rule" = $3, "b_rule" = $4
	       WHERE "id" = $5 AND "account_id" = $6
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q,
		entity.LinkID, entity.URL,
		encodedRule, encodedBinaryRule,
		entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the target %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO issue#131
func (scope targetScope) Delete(token *types.Token, data query.DeleteTarget) (types.Target, error) {
	entity, readErr := scope.Read(token, query.ReadTarget{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.Permanently {
		q := `DELETE FROM "target" WHERE "id" = $1 AND "account_id" = $2 RETURNING now()`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
		if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
			return entity, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to delete the target %q permanently",
				token.UserID, token.User.AccountID, entity.ID)
		}
		return entity, nil
	}
	q := `UPDATE "target"
	         SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the target %q safely",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
