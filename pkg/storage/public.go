package storage

import (
	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/storage/postgres"
)

// Link returns the Link with its Aliases and Targets by provided ID.
func (storage *Storage) Link(id domain.ID) (domain.Link, error) {
	return postgres.Link(storage.db, id)
}

// LinkByAlias returns the Link with its set of Alias and set of Target defined by provided namespace and URN.
func (storage *Storage) LinkByAlias(ns, urn string) (domain.Link, error) {
	return postgres.LinkByAlias(storage.db, ns, urn)
}

// Log stores a "redirect event".
func (storage *Storage) Log(event domain.Log) (domain.Log, error) {
	return postgres.Log(storage.db, event)
}

// UUID returns a new generated unique identifier.
func (storage *Storage) UUID() (domain.ID, error) {
	return postgres.UUID(storage.db)
}
