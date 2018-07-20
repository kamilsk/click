package grpc

import (
	"context"
	"log"
)

// NewTargetServer returns new instance of server API for Target service.
func NewTargetServer() TargetServer {
	return &targetServer{}
}

type targetServer struct {
}

func (*targetServer) Create(context.Context, *CreateTargetRequest) (*CreateTargetResponse, error) {
	log.Println("TargetServer.Create was called")
	return &CreateTargetResponse{}, nil
}

func (*targetServer) Read(context.Context, *ReadTargetRequest) (*ReadTargetResponse, error) {
	log.Println("TargetServer.Read was called")
	return &ReadTargetResponse{}, nil
}

func (*targetServer) Update(context.Context, *UpdateTargetRequest) (*UpdateTargetResponse, error) {
	log.Println("TargetServer.Update was called")
	return &UpdateTargetResponse{}, nil
}

func (*targetServer) Delete(context.Context, *DeleteTargetRequest) (*DeleteTargetResponse, error) {
	log.Println("TargetServer.Delete was called")
	return &DeleteTargetResponse{}, nil
}
