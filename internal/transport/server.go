package transport

import "net"

// Server TODO issue#131
type Server interface {
	// Serve TODO issue#131
	Serve(net.Listener) error
}
