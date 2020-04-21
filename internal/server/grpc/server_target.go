package grpc

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.octolab.org/ecosystem/click/internal/domain"
	"go.octolab.org/ecosystem/click/internal/server/grpc/middleware"
	"go.octolab.org/ecosystem/click/internal/storage/query"
)

// NewTargetServer returns new instance of server API for Target service.
func NewTargetServer(storage ProtectedStorage) TargetServer {
	return &targetServer{storage}
}

type targetServer struct {
	storage ProtectedStorage
}

// Create TODO issue#131
func (server *targetServer) Create(ctx context.Context, req *CreateTargetRequest) (*CreateTargetResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	var rule domain.Rule
	if len(req.Rule) > 0 {
		if err := json.Unmarshal([]byte(req.Rule), &rule); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid rule provided: %+v", err)
		}
	}
	target, createErr := server.storage.CreateTarget(ctx, token, query.CreateTarget{
		ID:         ptrToID(req.Id),
		LinkID:     domain.ID(req.LinkId),
		URL:        req.Url,
		Rule:       rule,
		BinaryRule: domain.BinaryRule(req.BRule),
	})
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", createErr)
	}
	return &CreateTargetResponse{
		Id:        target.ID.String(),
		CreatedAt: Timestamp(&target.CreatedAt),
	}, nil
}

// Read TODO issue#131
func (server *targetServer) Read(ctx context.Context, req *ReadTargetRequest) (*ReadTargetResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	target, readErr := server.storage.ReadTarget(ctx, token, query.ReadTarget{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
	}
	rule, _ := json.Marshal(target.Rule)
	return &ReadTargetResponse{
		Id:        target.ID.String(),
		LinkId:    target.LinkID.String(),
		Url:       target.URL,
		Rule:      string(rule),
		BRule:     string(target.BinaryRule),
		CreatedAt: Timestamp(&target.CreatedAt),
		UpdatedAt: Timestamp(target.UpdatedAt),
		DeletedAt: Timestamp(target.DeletedAt),
	}, nil
}

// Update TODO issue#131
func (server *targetServer) Update(ctx context.Context, req *UpdateTargetRequest) (*UpdateTargetResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	var rule domain.Rule
	if len(req.Rule) > 0 {
		if err := json.Unmarshal([]byte(req.Rule), &rule); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid rule provided: %+v", err)
		}
	}
	target, updateErr := server.storage.UpdateTarget(ctx, token, query.UpdateTarget{
		ID:         domain.ID(req.Id),
		LinkID:     domain.ID(req.LinkId),
		URL:        req.Url,
		Rule:       rule,
		BinaryRule: domain.BinaryRule(req.BRule),
	})
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", updateErr)
	}
	return &UpdateTargetResponse{
		Id:        target.ID.String(),
		UpdatedAt: Timestamp(target.UpdatedAt),
	}, nil
}

// Delete TODO issue#131
func (server *targetServer) Delete(ctx context.Context, req *DeleteTargetRequest) (*DeleteTargetResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	target, deleteErr := server.storage.DeleteTarget(ctx, token, query.DeleteTarget{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", deleteErr)
	}
	return &DeleteTargetResponse{
		Id:        target.ID.String(),
		DeletedAt: Timestamp(target.DeletedAt),
	}, nil
}
