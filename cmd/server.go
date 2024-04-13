package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shikachuu/template-files/internal"
	"github.com/Shikachuu/template-files/pkg"
	"github.com/Shikachuu/template-files/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterTemplateServiceServer(s, pkg.NewServer(&internal.DummyDatabase{}))

	log.Println("starting server...")
	go (func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	})()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down server...")
	s.GracefulStop()
}
