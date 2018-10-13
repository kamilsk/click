package grpc

import (
	"context"
	"log"
)

// NewEventServer returns new instance of server API for Log service.
func NewEventServer(storage ProtectedStorage) EventServer {
	return &logServer{storage}
}

type logServer struct {
	storage ProtectedStorage
}

// Read TODO issue#131
func (*logServer) Read(context.Context, *ReadEventsRequest) (*ReadEventsResponse, error) {
	log.Println("EventServer.Read was called")
	return &ReadEventsResponse{}, nil
}

// Listen TODO issue#131
func (*logServer) Listen(*ListenEventsRequest, Event_ListenServer) error {
	log.Println("EventServer.Listen was called")
	return nil
}
