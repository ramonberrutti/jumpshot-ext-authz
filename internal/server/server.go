package server

import (
	"context"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ramonberrutti/jumpshot-ext-authz/internal/server/auth"
	"github.com/ramonberrutti/jumpshot-ext-authz/internal/storage"
	"google.golang.org/grpc"

	authpb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
)

type Server interface {
	Run(ctx context.Context) error
}

type Options func(*server)

func WithListener(listener net.Listener) func(*server) {
	return func(o *server) {
		o.listener = listener
	}
}

func WithMetricsListener(listener net.Listener) func(*server) {
	return func(o *server) {
		o.metricsListener = listener
	}
}

func WithAuthStorage(store storage.Storage) func(*server) {
	return func(o *server) {
		o.authStore = store
	}
}

func NewServer(opts ...Options) Server {
	s := &server{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type server struct {
	listener        net.Listener
	metricsListener net.Listener

	authStore storage.Storage
}

func (s *server) Run(ctx context.Context) error {
	metrics := grpc_prometheus.NewServerMetrics(grpc_prometheus.WithServerHandlingTimeHistogram())
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			metrics.UnaryServerInterceptor(),
		),
	)

	authpb.RegisterAuthorizationServer(srv, &auth.AuthorizationServiceServer{
		Store: s.authStore,
	})

	if err := prometheus.Register(metrics); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	httpServer := &http.Server{
		Handler: mux,
	}

	errChan := make(chan error, 1)

	go func() {
		errChan <- srv.Serve(s.listener)
	}()

	if s.metricsListener != nil {
		go func() {
			errChan <- httpServer.Serve(s.metricsListener)
		}()
	}

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
	}

	srv.GracefulStop()
	return httpServer.Shutdown(ctx)
}
