package grpc

import (
	"context"
	"log"
)

// NewNamespaceServer returns new instance of server API for Namespace service.
func NewNamespaceServer() NamespaceServer {
	return &namespaceServer{}
}

type namespaceServer struct {
}

func (*namespaceServer) Create(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	log.Println("NamespaceServer.Create was called")
	return &CreateNamespaceResponse{}, nil
}

func (*namespaceServer) Read(context.Context, *ReadNamespaceRequest) (*ReadNamespaceResponse, error) {
	log.Println("NamespaceServer.Read was called")
	return &ReadNamespaceResponse{}, nil
}

func (*namespaceServer) Update(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	log.Println("NamespaceServer.Update was called")
	return &UpdateNamespaceResponse{}, nil
}

func (*namespaceServer) Delete(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	log.Println("NamespaceServer.Delete was called")
	return &DeleteNamespaceResponse{}, nil
}
