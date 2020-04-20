package storage

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/kamilsk/click/pkg/storage/types"
)

/*
 *
 * Link
 *
 */

// CreateLink TODO issue#131
func (storage *Storage) CreateLink(ctx context.Context, tokenID domain.ID, data query.CreateLink) (types.Link, error) {
	var entity types.Link

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.LinkEditor(ctx, conn).Create(token, data)
}

// ReadLink TODO issue#131
func (storage *Storage) ReadLink(ctx context.Context, tokenID domain.ID, data query.ReadLink) (types.Link, error) {
	var entity types.Link

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.LinkEditor(ctx, conn).Read(token, data)
}

// UpdateLink TODO issue#131
func (storage *Storage) UpdateLink(ctx context.Context, tokenID domain.ID, data query.UpdateLink) (types.Link, error) {
	var entity types.Link

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.LinkEditor(ctx, conn).Update(token, data)
}

// DeleteLink TODO issue#131
func (storage *Storage) DeleteLink(ctx context.Context, tokenID domain.ID, data query.DeleteLink) (types.Link, error) {
	var entity types.Link

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.LinkEditor(ctx, conn).Delete(token, data)
}

/*
 *
 * Namespace
 *
 */

// CreateNamespace TODO issue#131
func (storage *Storage) CreateNamespace(ctx context.Context, tokenID domain.ID, data query.CreateNamespace) (types.Namespace, error) {
	var entity types.Namespace

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.NamespaceEditor(ctx, conn).Create(token, data)
}

// ReadNamespace TODO issue#131
func (storage *Storage) ReadNamespace(ctx context.Context, tokenID domain.ID, data query.ReadNamespace) (types.Namespace, error) {
	var entity types.Namespace

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.NamespaceEditor(ctx, conn).Read(token, data)
}

// UpdateNamespace TODO issue#131
func (storage *Storage) UpdateNamespace(ctx context.Context, tokenID domain.ID, data query.UpdateNamespace) (types.Namespace, error) {
	var entity types.Namespace

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.NamespaceEditor(ctx, conn).Update(token, data)
}

// DeleteNamespace TODO issue#131
func (storage *Storage) DeleteNamespace(ctx context.Context, tokenID domain.ID, data query.DeleteNamespace) (types.Namespace, error) {
	var entity types.Namespace

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.NamespaceEditor(ctx, conn).Delete(token, data)
}

/*
 *
 * Alias
 *
 */

// CreateAlias TODO issue#131
func (storage *Storage) CreateAlias(ctx context.Context, tokenID domain.ID, data query.CreateAlias) (types.Alias, error) {
	var entity types.Alias

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	g := &errgroup.Group{}
	g.Go(func() error {
		_, readErr := storage.exec.LinkEditor(ctx, conn).Read(token, query.ReadLink{ID: data.LinkID})
		return readErr
	})
	g.Go(func() error {
		_, readErr := storage.exec.NamespaceEditor(ctx, conn).Read(token, query.ReadNamespace{ID: data.NamespaceID})
		return readErr
	})
	if readErr := g.Wait(); readErr != nil {
		return entity, readErr
	}
	return storage.exec.AliasEditor(ctx, conn).Create(token, data)
}

// ReadAlias TODO issue#131
func (storage *Storage) ReadAlias(ctx context.Context, tokenID domain.ID, data query.ReadAlias) (types.Alias, error) {
	var entity types.Alias

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.AliasEditor(ctx, conn).Read(token, data)
}

// UpdateAlias TODO issue#131
func (storage *Storage) UpdateAlias(ctx context.Context, tokenID domain.ID, data query.UpdateAlias) (types.Alias, error) {
	var entity types.Alias

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	g := &errgroup.Group{}
	g.Go(func() error {
		_, readErr := storage.exec.LinkEditor(ctx, conn).Read(token, query.ReadLink{ID: data.LinkID})
		return readErr
	})
	g.Go(func() error {
		_, readErr := storage.exec.NamespaceEditor(ctx, conn).Read(token, query.ReadNamespace{ID: data.NamespaceID})
		return readErr
	})
	if readErr := g.Wait(); readErr != nil {
		return entity, readErr
	}
	return storage.exec.AliasEditor(ctx, conn).Update(token, data)
}

// DeleteAlias TODO issue#131
func (storage *Storage) DeleteAlias(ctx context.Context, tokenID domain.ID, data query.DeleteAlias) (types.Alias, error) {
	var entity types.Alias

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.AliasEditor(ctx, conn).Delete(token, data)
}

/*
 *
 * Target
 *
 */

// CreateTarget TODO issue#131
func (storage *Storage) CreateTarget(ctx context.Context, tokenID domain.ID, data query.CreateTarget) (types.Target, error) {
	var entity types.Target

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	if _, readErr := storage.exec.LinkEditor(ctx, conn).Read(token, query.ReadLink{ID: data.LinkID}); readErr != nil {
		return entity, readErr
	}
	return storage.exec.TargetEditor(ctx, conn).Create(token, data)
}

// ReadTarget TODO issue#131
func (storage *Storage) ReadTarget(ctx context.Context, tokenID domain.ID, data query.ReadTarget) (types.Target, error) {
	var entity types.Target

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.TargetEditor(ctx, conn).Read(token, data)
}

// UpdateTarget TODO issue#131
func (storage *Storage) UpdateTarget(ctx context.Context, tokenID domain.ID, data query.UpdateTarget) (types.Target, error) {
	var entity types.Target

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	if _, readErr := storage.exec.LinkEditor(ctx, conn).Read(token, query.ReadLink{ID: data.LinkID}); readErr != nil {
		return entity, readErr
	}
	return storage.exec.TargetEditor(ctx, conn).Update(token, data)
}

// DeleteTarget TODO issue#131
func (storage *Storage) DeleteTarget(ctx context.Context, tokenID domain.ID, data query.DeleteTarget) (types.Target, error) {
	var entity types.Target

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return entity, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if authErr != nil {
		return entity, authErr
	}
	return storage.exec.TargetEditor(ctx, conn).Delete(token, data)
}
