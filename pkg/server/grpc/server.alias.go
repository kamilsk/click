package grpc

import (
	"context"
	"log"
)

// NewAliasServer returns new instance of server API for Alias service.
func NewAliasServer() AliasServer {
	return &aliasServer{}
}

type aliasServer struct {
}

func (*aliasServer) Create(context.Context, *CreateAliasRequest) (*CreateAliasResponse, error) {
	log.Println("AliasServer.Create was called")
	return &CreateAliasResponse{}, nil
}

func (*aliasServer) Read(context.Context, *ReadAliasRequest) (*ReadAliasResponse, error) {
	log.Println("AliasServer.Read was called")
	return &ReadAliasResponse{}, nil
}

func (*aliasServer) Update(context.Context, *UpdateAliasRequest) (*UpdateAliasResponse, error) {
	log.Println("AliasServer.Update was called")
	return &UpdateAliasResponse{}, nil
}

func (*aliasServer) Delete(context.Context, *DeleteAliasRequest) (*DeleteAliasResponse, error) {
	log.Println("AliasServer.Delete was called")
	return &DeleteAliasResponse{}, nil
}
