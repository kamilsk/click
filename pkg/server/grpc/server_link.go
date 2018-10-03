package grpc

import (
	"context"
	"log"
)

// NewLinkServer returns new instance of server API for Link service.
func NewLinkServer(storage ProtectedStorage) LinkServer {
	return &linkServer{storage}
}

type linkServer struct {
	storage ProtectedStorage
}

// Create TODO issue#131
func (*linkServer) Create(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error) {
	log.Println("LinkServer.Create was called")
	return &CreateLinkResponse{}, nil
}

// Read TODO issue#131
func (*linkServer) Read(context.Context, *ReadLinkRequest) (*ReadLinkResponse, error) {
	log.Println("LinkServer.Read was called")
	return &ReadLinkResponse{}, nil
}

// Update TODO issue#131
func (*linkServer) Update(context.Context, *UpdateLinkRequest) (*UpdateLinkResponse, error) {
	log.Println("LinkServer.Update was called")
	return &UpdateLinkResponse{}, nil
}

// Delete TODO issue#131
func (*linkServer) Delete(context.Context, *DeleteLinkRequest) (*DeleteLinkResponse, error) {
	log.Println("LinkServer.Delete was called")
	return &DeleteLinkResponse{}, nil
}
