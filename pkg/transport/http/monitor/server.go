package monitor

import (
	"expvar"
	"net"
	"net/http"

	"github.com/kamilsk/click/pkg/config"
	"github.com/kamilsk/click/pkg/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// New TODO issue#131
func New(_ config.MonitoringConfig) transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#131
func (*server) Serve(listener net.Listener) error {
	defer listener.Close()
	mux := &http.ServeMux{}
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/vars", expvar.Handler())
	return http.Serve(listener, mux)
}
