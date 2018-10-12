package storage

import (
	"context"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/storage/query"
	"golang.org/x/sync/errgroup"
)

// Link returns the Link with its Aliases and Targets by provided ID.
func (storage *Storage) Link(ctx context.Context, id domain.ID) (domain.Link, error) {
	var link domain.Link

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return link, connErr
	}
	defer closer()

	g := &errgroup.Group{}
	g.Go(func() error {
		entity, err := storage.exec.LinkReader(ctx, conn).ReadByID(id)
		if err != nil {
			return err
		}
		link.ID, link.Name = entity.ID, entity.Name
		return nil
	})
	g.Go(func() error {
		entities, err := storage.exec.AliasReader(ctx, conn).ReadAllByLink(id)
		if err != nil {
			return err
		}
		link.Aliases = make(domain.Aliases, 0, len(entities))
		for _, entity := range entities {
			link.Aliases = append(link.Aliases, domain.Alias{
				ID:        entity.ID,
				Namespace: entity.NamespaceID,
				URN:       entity.URN,
			})
		}
		return nil
	})
	g.Go(func() error {
		entities, err := storage.exec.TargetReader(ctx, conn).ReadAllByLink(id)
		if err != nil {
			return err
		}
		link.Targets = make(domain.Targets, 0, len(entities))
		for _, entity := range entities {
			link.Targets = append(link.Targets, domain.Target{
				ID:   entity.ID,
				Rule: entity.Rule,
				URL:  entity.URI,
			})
		}
		return nil
	})
	return link, g.Wait()
}

// LinkByAlias returns the Link with its set of Alias and set of Target defined by provided namespace and URN.
func (storage *Storage) LinkByAlias(ctx context.Context, ns domain.ID, urn string) (domain.Link, error) {
	var link domain.Link

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return link, connErr
	}
	defer closer()

	entity, readErr := storage.exec.LinkReader(ctx, conn).ReadByAlias(ns, urn)
	if readErr != nil {
		return link, readErr
	}
	return storage.Link(ctx, entity.ID)
}

// LogRedirect stores a redirect event.
func (storage *Storage) LogRedirect(ctx context.Context, event domain.RedirectEvent) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer closer()

	// TODO issue#51
	_, writeErr := storage.exec.LogWriter(ctx, conn).Write(query.WriteLog{
		LinkID:          event.LinkID,
		AliasID:         event.AliasID,
		TargetID:        event.TargetID,
		Identifier:      event.Identifier,
		URI:             event.URL,
		Code:            event.Code,
		RedirectContext: event.Context,
	})

	return writeErr
}
