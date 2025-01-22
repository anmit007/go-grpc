package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/anmit007/go-grpc/grpc-streaming/internal/streaming"
	"github.com/anmit007/go-grpc/grpc-streaming/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running app", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("closing server gracefully")
}
func run(ctx context.Context) error {
	grpcServer := grpc.NewServer()
	streamingService := streaming.Service{}
	proto.RegisterStreamingServiceServer(grpcServer, streamingService)
	const addr = ":50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen on address: %w", err)
		}
		slog.Info("STARTING SERVER ON", slog.String("ADDRESS", addr))
		if err := grpcServer.Serve(listen); err != nil {
			return fmt.Errorf("failed to serve grpc service :%w", err)
		}
		return nil
	})
	g.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()
		return nil
	})
	return g.Wait()
}
