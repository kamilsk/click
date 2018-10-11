package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/errors"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/kamilsk/click/pkg/storage/types"
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
		URI:        data.URI,
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
	q := `INSERT INTO "target" ("id", "account_id", "link_id", "uri", "rule", "b_rule")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3, $4, $5, $6)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID, entity.LinkID,
		entity.URI, encodedRule, encodedBinaryRule)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a target %q",
			token.UserID, token.User.AccountID, entity.URI)
	}
	return entity, nil
}

// Read TODO issue#131
func (scope targetScope) Read(token *types.Token, data query.ReadTarget) (types.Target, error) {
	var encodedRule, encodedBinaryRule []byte
	entity := types.Target{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "link_id", "uri", "rule", "b_rule", "created_at", "updated_at", "deleted_at"
	        FROM "target"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.LinkID, &entity.URI,
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
	q := `SELECT "id", "uri", "rule", "b_rule", "created_at", "updated_at", "deleted_at"
	        FROM "target"
	       WHERE "link_id" = $1 AND "deleted_at" IS NULL`
	rows, queryErr := scope.conn.QueryContext(scope.ctx, q, link)
	if queryErr != nil {
		return nil, errors.Database(errors.ServerErrorMessage, queryErr,
			"trying to read all targets of the link %q", link)
	}
	defer rows.Close()
	result := make([]types.Target, 0, 4)
	for rows.Next() {
		var encodedRule, encodedBinaryRule []byte
		entity := types.Target{LinkID: link}
		if scanErr := rows.Scan(&entity.ID, &entity.URI,
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
	return result, nil
}

// Update TODO issue#131
func (scope targetScope) Update(token *types.Token, data query.UpdateTarget) (types.Target, error) {
	entity, readErr := scope.Read(token, query.ReadTarget{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.URI != "" {
		entity.URI = data.URI
	}
	if !data.Rule.IsEmpty() {
		entity.Rule = data.Rule
	}
	if !data.BinaryRule.IsEmpty() {
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
	         SET "uri" = $1, "rule" = $2, "b_rule" = $3
	       WHERE "id" = $4 AND "account_id" = $5
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.URI, encodedRule, encodedBinaryRule,
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
