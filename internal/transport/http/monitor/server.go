package monitor

import (
	"expvar"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.octolab.org/ecosystem/click/internal/config"
	"go.octolab.org/ecosystem/click/internal/transport"
)

// New TODO issue#131
func New(_ config.MonitoringConfig) transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#131
func (*server) Serve(listener net.Listener) error {
	defer func() { _ = listener.Close() }()
	mux := &http.ServeMux{}
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/vars", expvar.Handler())
	return http.Serve(listener, mux)
}
