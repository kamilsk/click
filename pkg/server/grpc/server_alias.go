package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/server/grpc/middleware"
	"github.com/kamilsk/click/pkg/storage/query"
)

// NewAliasServer returns new instance of server API for Alias service.
func NewAliasServer(storage ProtectedStorage) AliasServer {
	return &aliasServer{storage}
}

type aliasServer struct {
	storage ProtectedStorage
}

// Create TODO issue#131
func (server *aliasServer) Create(ctx context.Context, req *CreateAliasRequest) (*CreateAliasResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	alias, createErr := server.storage.CreateAlias(ctx, token, query.CreateAlias{
		ID:          ptrToID(req.Id),
		LinkID:      domain.ID(req.LinkId),
		NamespaceID: domain.ID(req.NamespaceId),
		URN:         req.Urn,
	})
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", createErr)
	}
	return &CreateAliasResponse{
		Id:        alias.ID.String(),
		CreatedAt: Timestamp(&alias.CreatedAt),
	}, nil
}

// Read TODO issue#131
func (server *aliasServer) Read(ctx context.Context, req *ReadAliasRequest) (*ReadAliasResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	alias, readErr := server.storage.ReadAlias(ctx, token, query.ReadAlias{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
	}
	return &ReadAliasResponse{
		Id:          alias.ID.String(),
		LinkId:      alias.LinkID.String(),
		NamespaceId: alias.NamespaceID.String(),
		Urn:         alias.URN,
		CreatedAt:   Timestamp(&alias.CreatedAt),
		UpdatedAt:   Timestamp(alias.UpdatedAt),
		DeletedAt:   Timestamp(alias.DeletedAt),
	}, nil
}

// Update TODO issue#131
func (server *aliasServer) Update(ctx context.Context, req *UpdateAliasRequest) (*UpdateAliasResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	alias, updateErr := server.storage.UpdateAlias(ctx, token, query.UpdateAlias{
		ID:          domain.ID(req.Id),
		LinkID:      domain.ID(req.LinkId),
		NamespaceID: domain.ID(req.NamespaceId),
		URN:         req.Urn,
	})
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", updateErr)
	}
	return &UpdateAliasResponse{
		Id:        alias.ID.String(),
		UpdatedAt: Timestamp(alias.UpdatedAt),
	}, nil
}

// Delete TODO issue#131
func (server *aliasServer) Delete(ctx context.Context, req *DeleteAliasRequest) (*DeleteAliasResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	alias, deleteErr := server.storage.DeleteAlias(ctx, token, query.DeleteAlias{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", deleteErr)
	}
	return &DeleteAliasResponse{
		Id:        alias.ID.String(),
		DeletedAt: Timestamp(alias.DeletedAt),
	}, nil
}
