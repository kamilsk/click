package grpc

import (
	"context"
	"log"
)

// NewLinkServer returns new instance of server API for Link service.
func NewLinkServer() LinkServer {
	return &linkServer{}
}

type linkServer struct {
}

func (*linkServer) Create(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error) {
	log.Println("LinkServer.Create was called")
	return &CreateLinkResponse{}, nil
}

func (*linkServer) Read(context.Context, *ReadLinkRequest) (*ReadLinkResponse, error) {
	log.Println("LinkServer.Read was called")
	return &ReadLinkResponse{}, nil
}

func (*linkServer) Update(context.Context, *UpdateLinkRequest) (*UpdateLinkResponse, error) {
	log.Println("LinkServer.Update was called")
	return &UpdateLinkResponse{}, nil
}

func (*linkServer) Delete(context.Context, *DeleteLinkRequest) (*DeleteLinkResponse, error) {
	log.Println("LinkServer.Delete was called")
	return &DeleteLinkResponse{}, nil
}
