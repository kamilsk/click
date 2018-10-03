package grpc

import (
	"context"
	"log"
)

// NewNamespaceServer returns new instance of server API for Namespace service.
func NewNamespaceServer(storage ProtectedStorage) NamespaceServer {
	return &namespaceServer{storage}
}

type namespaceServer struct {
	storage ProtectedStorage
}

// Create TODO issue#131
func (*namespaceServer) Create(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	log.Println("NamespaceServer.Create was called")
	return &CreateNamespaceResponse{}, nil
}

// Read TODO issue#131
func (*namespaceServer) Read(context.Context, *ReadNamespaceRequest) (*ReadNamespaceResponse, error) {
	log.Println("NamespaceServer.Read was called")
	return &ReadNamespaceResponse{}, nil
}

// Update TODO issue#131
func (*namespaceServer) Update(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	log.Println("NamespaceServer.Update was called")
	return &UpdateNamespaceResponse{}, nil
}

// Delete TODO issue#131
func (*namespaceServer) Delete(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	log.Println("NamespaceServer.Delete was called")
	return &DeleteNamespaceResponse{}, nil
}
