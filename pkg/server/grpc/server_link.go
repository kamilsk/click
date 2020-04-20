package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/server/grpc/middleware"
	"github.com/kamilsk/click/pkg/storage/query"
)

// NewLinkServer returns new instance of server API for Link service.
func NewLinkServer(storage ProtectedStorage) LinkServer {
	return &linkServer{storage}
}

type linkServer struct {
	storage ProtectedStorage
}

// Create TODO issue#131
func (server *linkServer) Create(ctx context.Context, req *CreateLinkRequest) (*CreateLinkResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	link, createErr := server.storage.CreateLink(ctx, token, query.CreateLink{
		ID:   ptrToID(req.Id),
		Name: req.Name,
	})
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", createErr)
	}
	return &CreateLinkResponse{
		Id:        link.ID.String(),
		CreatedAt: Timestamp(&link.CreatedAt),
	}, nil
}

// Read TODO issue#131
func (server *linkServer) Read(ctx context.Context, req *ReadLinkRequest) (*ReadLinkResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	link, readErr := server.storage.ReadLink(ctx, token, query.ReadLink{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
	}
	return &ReadLinkResponse{
		Id:        link.ID.String(),
		Name:      link.Name,
		CreatedAt: Timestamp(&link.CreatedAt),
		UpdatedAt: Timestamp(link.UpdatedAt),
		DeletedAt: Timestamp(link.DeletedAt),
	}, nil
}

// Update TODO issue#131
func (server *linkServer) Update(ctx context.Context, req *UpdateLinkRequest) (*UpdateLinkResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	link, updateErr := server.storage.UpdateLink(ctx, token, query.UpdateLink{
		ID:   domain.ID(req.Id),
		Name: req.Name,
	})
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", updateErr)
	}
	return &UpdateLinkResponse{
		Id:        link.ID.String(),
		UpdatedAt: Timestamp(link.UpdatedAt),
	}, nil
}

// Delete TODO issue#131
func (server *linkServer) Delete(ctx context.Context, req *DeleteLinkRequest) (*DeleteLinkResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	link, deleteErr := server.storage.DeleteLink(ctx, token, query.DeleteLink{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", deleteErr)
	}
	return &DeleteLinkResponse{
		Id:        link.ID.String(),
		DeletedAt: Timestamp(link.DeletedAt),
	}, nil
}
