package query

import "github.com/kamilsk/click/pkg/domain"

// CreateNamespace TODO issue#131
type CreateNamespace struct {
	ID *domain.ID
}

// ReadNamespace TODO issue#131
type ReadNamespace struct {
	ID domain.ID
}

// UpdateNamespace TODO issue#131
type UpdateNamespace struct {
	ID domain.ID
}

// DeleteNamespace TODO issue#131
type DeleteNamespace struct {
	ID          domain.ID
	Permanently bool
}
