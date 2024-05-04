package main

import (
	"context"
	"log"
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
	"google.golang.org/grpc"
)

func main() {
    ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}

func run(ctx context.Context) error {
	var (
		db internal.DummyDatabase
		wg sync.WaitGroup
	)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg.Add(1)
	go runGRPCServer(&wg, ctx, &db, "8080")
    
    wg.Add(1)
    go runHTTPServer(&wg, ctx, &db, "8081")
	
    wg.Wait()
	return nil
}

func runHTTPServer(wg *sync.WaitGroup, ctx context.Context, db pkg.Database, port string) {
	httpServer := &http.Server{
		Addr:         net.JoinHostPort("", port),
		Handler:      web.NewHTTPHandler(db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		log.Println("starting http server...")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		<-ctx.Done()
		log.Println("shutting down http server...")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown http server: %v", err)
		}
	}()
}

func runGRPCServer(wg *sync.WaitGroup, ctx context.Context, db pkg.Database, port string) {
	defer wg.Done()

	lis, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterTemplateServiceServer(s, pkg.NewServer(db))

	log.Println("starting grpc server...")
	go (func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	})()

	<-ctx.Done()
	log.Println("shutting down grpc server...")
	s.GracefulStop()
}
