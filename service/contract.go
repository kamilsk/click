package service

import "github.com/kamilsk/click/domain"

// Storage defines the behavior of Data Access Object.
type Storage interface {
	// Link returns the Link with its set of Alias and set of Target by provided ID.
	Link(domain.UUID) (domain.Link, error)
	// LinkByAlias returns the Link with its set of Alias and set of Target defined by provided namespace and URN.
	LinkByAlias(namespace, urn string) (domain.Link, error)
	// Log stores a "redirect event".
	Log(event domain.Log) (domain.Log, error)
	// UUID returns a new generated unique identifier.
	UUID() (domain.UUID, error)
}
