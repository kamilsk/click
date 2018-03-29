package service

import "github.com/kamilsk/click/domain"

// Storage defines the behavior of Data Access Object.
type Storage interface {
	// Link returns the Link with its Aliases and Targets by provided ID.
	Link(domain.UUID) (domain.Link, error)
	// LinkByAlias returns the Link with its Targets and the single Alias defined by Namespace and URN.
	LinkByAlias(domain.Alias) (domain.Link, error)
}
