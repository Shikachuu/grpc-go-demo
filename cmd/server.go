package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shikachuu/template-files/internal"
	"github.com/Shikachuu/template-files/pkg"
	"github.com/Shikachuu/template-files/pkg/web"
	"github.com/Shikachuu/template-files/proto"
	"go.etcd.io/bbolt"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stderr); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}

func run(ctx context.Context, w io.Writer) error {
	var wg sync.WaitGroup
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

    bb, err := bbolt.Open("templates.db", 0600, nil)
    if err != nil {
        return err
    }
    defer bb.Close()

    db := internal.NewBBoltDatabase(bb)

	logger := slog.New(slog.NewTextHandler(
		w,
		&slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	))

	wg.Add(1)
	go runGRPCServer(&wg, ctx, logger.With("server", "grpc"), db, "8080")

	wg.Add(1)
	go runHTTPServer(&wg, ctx, logger.With("server", "http"), db, "8081")

	wg.Wait()
	return nil
}

func runHTTPServer(wg *sync.WaitGroup, ctx context.Context, logger *slog.Logger, db pkg.Database, port string) {
	httpServer := &http.Server{
		Addr:         net.JoinHostPort("", port),
		Handler:      web.NewHTTPHandler(db, logger),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		logger.InfoContext(ctx, "starting http server...", "port", port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorContext(ctx, "failed to serve http", "error", err.Error())
			os.Exit(1)
		}
	}()

	go func() {
		defer wg.Done()

		<-ctx.Done()
		logger.InfoContext(ctx, "shutting down http server...")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.ErrorContext(ctx, "failed to shutdown http server", "error", err.Error())
			os.Exit(1)
		}
	}()
}

func runGRPCServer(wg *sync.WaitGroup, ctx context.Context, logger *slog.Logger, db pkg.Database, port string) {
	defer wg.Done()

	lis, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		logger.ErrorContext(ctx, "failed to listen", "error", err.Error())
		os.Exit(1)
	}

	s := grpc.NewServer()
	proto.RegisterTemplateServiceServer(s, pkg.NewServer(db, logger))

	logger.InfoContext(ctx, "starting grpc server...", "port", port)
	go (func() {
		if err := s.Serve(lis); err != nil {
			logger.ErrorContext(ctx, "failed to serve grpc", "error", err.Error())
			os.Exit(1)
		}
	})()

	<-ctx.Done()
	logger.InfoContext(ctx, "shutting down grpc server...")
	s.GracefulStop()
}
