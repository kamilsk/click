package grpc

import (
	"context"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/storage/query"
	"github.com/kamilsk/click/pkg/storage/types"
)

// ProtectedStorage TODO issue#131
type ProtectedStorage interface {
	// CreateLink TODO issue#131
	CreateLink(ctx context.Context, tokenID domain.ID, data query.CreateLink) (types.Link, error)
	// ReadLink TODO issue#131
	ReadLink(ctx context.Context, tokenID domain.ID, data query.ReadLink) (types.Link, error)
	// UpdateLink TODO issue#131
	UpdateLink(ctx context.Context, tokenID domain.ID, data query.UpdateLink) (types.Link, error)
	// DeleteLink TODO issue#131
	DeleteLink(ctx context.Context, tokenID domain.ID, data query.DeleteLink) (types.Link, error)

	// CreateTarget TODO issue#131
	CreateTarget(ctx context.Context, tokenID domain.ID, data query.CreateTarget) (types.Target, error)
	// ReadTarget TODO issue#131
	ReadTarget(ctx context.Context, tokenID domain.ID, data query.ReadTarget) (types.Target, error)
	// UpdateTarget TODO issue#131
	UpdateTarget(ctx context.Context, tokenID domain.ID, data query.UpdateTarget) (types.Target, error)
	// DeleteTarget TODO issue#131
	DeleteTarget(ctx context.Context, tokenID domain.ID, data query.DeleteTarget) (types.Target, error)

	// CreateNamespace TODO issue#131
	CreateNamespace(ctx context.Context, tokenID domain.ID, data query.CreateNamespace) (types.Namespace, error)
	// ReadNamespace TODO issue#131
	ReadNamespace(ctx context.Context, tokenID domain.ID, data query.ReadNamespace) (types.Namespace, error)
	// UpdateNamespace TODO issue#131
	UpdateNamespace(ctx context.Context, tokenID domain.ID, data query.UpdateNamespace) (types.Namespace, error)
	// DeleteNamespace TODO issue#131
	DeleteNamespace(ctx context.Context, tokenID domain.ID, data query.DeleteNamespace) (types.Namespace, error)

	// CreateAlias TODO issue#131
	CreateAlias(ctx context.Context, tokenID domain.ID, data query.CreateAlias) (types.Alias, error)
	// ReadAlias TODO issue#131
	ReadAlias(ctx context.Context, tokenID domain.ID, data query.ReadAlias) (types.Alias, error)
	// UpdateAlias TODO issue#131
	UpdateAlias(ctx context.Context, tokenID domain.ID, data query.UpdateAlias) (types.Alias, error)
	// DeleteAlias TODO issue#131
	DeleteAlias(ctx context.Context, tokenID domain.ID, data query.DeleteAlias) (types.Alias, error)
}
