package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ekkinox/otlp-log-processor/internal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	lis, err := net.Listen("tcp", ":4317")
	if err != nil {
		log.Fatal(err)
	}

	srv := internal.NewServer()

	go func() {
		fmt.Println("starting server on :4317...")
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("stopping server...")
	srv.GracefulStop()
	log.Println("server stopped")
}
