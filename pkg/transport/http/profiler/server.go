package profiler

import (
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/kamilsk/click/pkg/config"
	"github.com/kamilsk/click/pkg/transport"
)

// New TODO issue#131
func New(_ config.ProfilingConfig) transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#131
func (*server) Serve(listener net.Listener) error {
	defer listener.Close()
	mux := &http.ServeMux{}
	mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/pprof/profile", pprof.Profile)
	mux.HandleFunc("/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	return http.Serve(listener, mux)
}
