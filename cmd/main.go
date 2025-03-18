package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ekkinox/otlp-log-processor/internal"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	stopOTel, err := internal.SetupOTel(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stopApp := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stopApp()

	lis, err := net.Listen("tcp", ":4317")
	if err != nil {
		log.Fatal(err)
	}

	srv := internal.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	go func() {
		fmt.Println("starting server on :4317...")
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("stopping server...")
	srv.GracefulStop()
	log.Println("stopping OTel...")
	stopOTel()
	log.Println("stopped")
}
