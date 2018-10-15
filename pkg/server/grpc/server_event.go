package grpc

import (
	"context"
	"log"
)

// NewEventServer returns new instance of server API for Tracker service.
func NewEventServer(storage ProtectedStorage) EventServer {
	return &eventServer{storage}
}

type eventServer struct {
	storage ProtectedStorage
}

// Read TODO issue#131
func (*eventServer) Read(context.Context, *ReadEventsRequest) (*ReadEventsResponse, error) {
	log.Println("EventServer.Read was called")
	return &ReadEventsResponse{}, nil
}

// Listen TODO issue#131
func (*eventServer) Listen(*ListenEventsRequest, Event_ListenServer) error {
	log.Println("EventServer.Listen was called")
	return nil
}
