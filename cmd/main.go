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

	fmt.Println("starting...")

	cfg := internal.NewConfig()
	str := internal.NewStorage()
	svc := internal.NewService(str, cfg.Attribute(), cfg.Workers())
	srv := internal.NewServer(svc, grpc.StatsHandler(otelgrpc.NewServerHandler()))
	tkr := internal.NewTicker(os.Stdout, str, cfg.Attribute(), cfg.Interval())

	fmt.Println("starting OTel components...")
	stopOTel, err := internal.SetupOTel(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)

	go func() {
		lis, err := net.Listen("tcp", ":4317")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("starting gRPC server on :4317...")
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		fmt.Println("starting ticker...")
		tkr.Start(ctx)
	}()

	<-ctx.Done()

	log.Println("stopping gRPC server...")
	srv.GracefulStop()

	log.Println("stopping OTel components...")
	stopOTel()

	log.Println("stopping...")
	stop()

	log.Println("stopped")
}
