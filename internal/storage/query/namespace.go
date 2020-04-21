package query

import "go.octolab.org/ecosystem/click/internal/domain"

// CreateNamespace TODO issue#131
type CreateNamespace struct {
	ID   *domain.ID
	Name string
}

// ReadNamespace TODO issue#131
type ReadNamespace struct {
	ID domain.ID
}

// UpdateNamespace TODO issue#131
type UpdateNamespace struct {
	ID   domain.ID
	Name string
}

// DeleteNamespace TODO issue#131
type DeleteNamespace struct {
	ID          domain.ID
	Permanently bool
}
