package grpc

import (
	"context"
	"log"
)

// NewEventServer returns new instance of server API for Tracker service.
func NewEventServer(storage ProtectedStorage) ListenerServer {
	return &eventServer{storage}
}

type eventServer struct {
	storage ProtectedStorage
}

// Read TODO issue#131
func (*eventServer) Read(context.Context, *ReadEventsRequest) (*ReadEventsResponse, error) {
	log.Println("ListenerServer.Read was called")
	return &ReadEventsResponse{}, nil
}

// Listen TODO issue#131
func (*eventServer) Listen(*ListenEventsRequest, Listener_ListenServer) error {
	log.Println("ListenerServer.Listen was called")
	return nil
}
