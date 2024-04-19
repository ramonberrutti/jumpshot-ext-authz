package main

import (
	"context"
	"fmt"
	"net"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/ramonberrutti/jumpshot-ext-authz/internal/server"
	"github.com/ramonberrutti/jumpshot-ext-authz/internal/storage"
	"github.com/ramonberrutti/jumpshot-ext-authz/internal/storage/engine"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type serveFlags struct {
	listen  string
	metrics string
	nats    string
}

func serveCmd() *cobra.Command {
	var flags serveFlags
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runServe(cmd.Context(), flags)
		},
	}

	serveCmd.Flags().StringVarP(&flags.listen, "listen", "l", ":9090", "Listen address")
	serveCmd.Flags().StringVar(&flags.metrics, "metrics", ":9091", "Metrics address")
	serveCmd.Flags().StringVarP(&flags.nats, "nats", "n", "nats://nats:4222", "NATS address")
	return serveCmd
}

func runServe(ctx context.Context, flags serveFlags) error {
	listener, err := net.Listen("tcp", flags.listen)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", flags.listen, err)
	}
	zap.L().Info("Listening on", zap.String("address", flags.listen))

	metricListener, err := net.Listen("tcp", flags.metrics)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", flags.metrics, err)
	}
	zap.L().Info("Metrics Listening on", zap.String("address", flags.metrics))

	nc, err := nats.Connect(flags.nats, nats.MaxReconnects(-1))
	if err != nil {
		return err
	}
	defer nc.Close()
	zap.L().Info("Connected to NATS", zap.String("address", flags.nats))

	js, err := jetstream.New(nc)
	if err != nil {
		return err
	}

	natsStorage := storage.NewStorage(engine.NewJetStream(js))

	srv := server.NewServer(
		server.WithListener(listener),
		server.WithMetricsListener(metricListener),
		server.WithAuthStorage(natsStorage),
	)

	return srv.Run(ctx)
}
