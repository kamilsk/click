package cmd

import (
	"expvar"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"runtime"
	"time"

	pb "github.com/kamilsk/click/pkg/server/grpc"

	"github.com/kamilsk/click/pkg/config"
	"github.com/kamilsk/click/pkg/server"
	"github.com/kamilsk/click/pkg/server/router/chi"
	"github.com/kamilsk/click/pkg/service"
	"github.com/kamilsk/click/pkg/storage"
	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime.GOMAXPROCS(int(cnf.Union.ServerConfig.CPUCount))

		if err := startGRPCServer(cnf.Union.GRPCConfig); err != nil {
			return err
		}
		if cnf.Union.MonitoringConfig.Enabled {
			if err := startMonitoring(cnf.Union.MonitoringConfig); err != nil {
				return err
			}
		}
		if cnf.Union.ProfilingConfig.Enabled {
			if err := startProfiler(cnf.Union.ProfilingConfig); err != nil {
				return err
			}
		}

		handler := chi.NewRouter(
			server.New(
				service.New(
					storage.Must(storage.Connection(cnf.Union.DBConfig)),
				),
			),
		)
		return startHTTPServer(cnf.Union.ServerConfig, handler)
	},
}

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("max_cpus") },
		func() error { return v.BindEnv("host") },
		func() error { return v.BindEnv("read_timeout") },
		func() error { return v.BindEnv("read_header_timeout") },
		func() error { return v.BindEnv("write_timeout") },
		func() error { return v.BindEnv("idle_timeout") },
		func() error { return v.BindEnv("profiling_host") },
		func() error { return v.BindEnv("monitoring_host") },
		func() error { return v.BindEnv("grpc_host") },
		func() error {
			v.SetDefault("max_cpus", 1)
			v.SetDefault("host", "127.0.0.1:80")
			v.SetDefault("read_timeout", time.Duration(0))
			v.SetDefault("read_header_timeout", time.Duration(0))
			v.SetDefault("write_timeout", time.Duration(0))
			v.SetDefault("idle_timeout", time.Duration(0))
			v.SetDefault("profiling_host", "127.0.0.1:8090")
			v.SetDefault("monitoring_host", "127.0.0.1:8091")
			v.SetDefault("grpc_host", "127.0.0.1:8092")
			return nil
		},
		func() error {
			flags := runCmd.Flags()
			flags.UintVarP(&cnf.Union.ServerConfig.CPUCount,
				"cpus", "C", uint(v.GetInt("max_cpus")), "maximum number of CPUs that can be executing simultaneously")
			flags.StringVarP(&cnf.Union.ServerConfig.Interface,
				"host", "H", v.GetString("host"), "web server host")
			flags.DurationVarP(&cnf.Union.ServerConfig.ReadTimeout,
				"read-timeout", "", v.GetDuration("read_timeout"),
				"maximum duration for reading the entire request, including the body")
			flags.DurationVarP(&cnf.Union.ServerConfig.ReadHeaderTimeout,
				"read-header-timeout", "", v.GetDuration("read_header_timeout"),
				"amount of time allowed to read request headers")
			flags.DurationVarP(&cnf.Union.ServerConfig.WriteTimeout,
				"write-timeout", "", v.GetDuration("write_timeout"),
				"maximum duration before timing out writes of the response")
			flags.DurationVarP(&cnf.Union.ServerConfig.IdleTimeout,
				"idle-timeout", "", v.GetDuration("idle_timeout"),
				"maximum amount of time to wait for the next request when keep-alive is enabled")
			flags.BoolVarP(&cnf.Union.ProfilingConfig.Enabled,
				"with-profiling", "", false, "enable pprof on /pprof/* and /debug/pprof/")
			flags.StringVarP(&cnf.Union.ProfilingConfig.Interface,
				"profiling-host", "", v.GetString("profiling_host"), "profiling host")
			flags.BoolVarP(&cnf.Union.MonitoringConfig.Enabled,
				"with-monitoring", "", false, "enable prometheus on /monitoring and expvar on /vars")
			flags.StringVarP(&cnf.Union.MonitoringConfig.Interface,
				"monitoring-host", "", v.GetString("monitoring_host"), "monitoring host")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			return nil
		},
	)
	db(runCmd)
}

func startHTTPServer(cnf config.ServerConfig, handler http.Handler) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	srv := &http.Server{Addr: cnf.Interface, Handler: handler,
		ReadTimeout:       cnf.ReadTimeout,
		ReadHeaderTimeout: cnf.ReadHeaderTimeout,
		WriteTimeout:      cnf.WriteTimeout,
		IdleTimeout:       cnf.IdleTimeout,
	}
	log.Println("start HTTP server at", listener.Addr())
	return srv.Serve(listener)
}

func startGRPCServer(cnf config.GRPCConfig) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	go func() {
		srv := grpc.NewServer()
		pb.RegisterNamespaceServer(srv, pb.NewNamespaceServer())
		pb.RegisterLinkServer(srv, pb.NewLinkServer())
		pb.RegisterAliasServer(srv, pb.NewAliasServer())
		pb.RegisterTargetServer(srv, pb.NewTargetServer())
		pb.RegisterLogServer(srv, pb.NewLogServer())
		log.Println("start gRPC server at", listener.Addr())
		_ = srv.Serve(listener) // TODO issue#97
		listener.Close()
	}()
	return nil
}

func startMonitoring(cnf config.MonitoringConfig) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	go func() {
		mux := &http.ServeMux{}
		mux.Handle("/monitoring", promhttp.Handler())
		mux.Handle("/vars", expvar.Handler())
		log.Println("start monitor at", listener.Addr())
		_ = http.Serve(listener, mux) // TODO issue#97
		listener.Close()
	}()
	return nil
}

func startProfiler(cnf config.ProfilingConfig) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	go func() {
		mux := &http.ServeMux{}
		mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/pprof/profile", pprof.Profile)
		mux.HandleFunc("/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/pprof/trace", pprof.Trace)
		mux.HandleFunc("/debug/pprof/", pprof.Index) // net/http/pprof.handler.ServeHTTP specificity
		log.Println("start profiler at", listener.Addr())
		_ = http.Serve(listener, mux) // TODO issue#97
		listener.Close()
	}()
	return nil
}
