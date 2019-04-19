package grpc

import (
	"context"
	"github.com/MaksimYudenko/training/finalTask/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

// RunServer runs gRPC service to publish Training service
func RunServer(
	ctx context.Context, v1API v1.TraineeServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// register service
	server := grpc.NewServer()
	v1.RegisterTraineeServiceServer(server, v1API)
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down the server ...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
