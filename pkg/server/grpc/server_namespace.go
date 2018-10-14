package grpc

import (
	"context"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/server/grpc/middleware"
	"github.com/kamilsk/click/pkg/storage/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewNamespaceServer returns new instance of server API for Namespace service.
func NewNamespaceServer(storage ProtectedStorage) NamespaceServer {
	return &namespaceServer{storage}
}

type namespaceServer struct {
	storage ProtectedStorage
}

// Create TODO issue#131
func (server *namespaceServer) Create(ctx context.Context, req *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	namespace, createErr := server.storage.CreateNamespace(ctx, token, query.CreateNamespace{
		ID:   ptrToID(req.Id),
		Name: req.Name,
	})
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", createErr)
	}
	return &CreateNamespaceResponse{
		Id:        namespace.ID.String(),
		CreatedAt: Timestamp(&namespace.CreatedAt),
	}, nil
}

// Read TODO issue#131
func (server *namespaceServer) Read(ctx context.Context, req *ReadNamespaceRequest) (*ReadNamespaceResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	namespace, readErr := server.storage.ReadNamespace(ctx, token, query.ReadNamespace{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
	}
	return &ReadNamespaceResponse{
		Id:        namespace.ID.String(),
		Name:      namespace.Name,
		CreatedAt: Timestamp(&namespace.CreatedAt),
		UpdatedAt: Timestamp(namespace.UpdatedAt),
		DeletedAt: Timestamp(namespace.DeletedAt),
	}, nil
}

// Update TODO issue#131
func (server *namespaceServer) Update(ctx context.Context, req *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	namespace, updateErr := server.storage.UpdateNamespace(ctx, token, query.UpdateNamespace{
		ID:   domain.ID(req.Id),
		Name: req.Name,
	})
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", updateErr)
	}
	return &UpdateNamespaceResponse{
		Id:        namespace.ID.String(),
		UpdatedAt: Timestamp(namespace.UpdatedAt),
	}, nil
}

// Delete TODO issue#131
func (server *namespaceServer) Delete(ctx context.Context, req *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	namespace, deleteErr := server.storage.DeleteNamespace(ctx, token, query.DeleteNamespace{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", deleteErr)
	}
	return &DeleteNamespaceResponse{
		Id:        namespace.ID.String(),
		DeletedAt: Timestamp(namespace.DeletedAt),
	}, nil
}
