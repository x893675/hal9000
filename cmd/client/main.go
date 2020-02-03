package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"hal9000/internal/addsrv"
	"log"
	"time"
)

func main() {
	gRPCAddr := flag.String("gRPC", ":8891", "gRPC client")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(),
		10000*time.Millisecond)
	defer cancel()

	conn, err := grpc.DialContext(ctx,
		*gRPCAddr, grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalln("gRPC dial error:", err)
	}
	defer conn.Close()

	addService := New(conn)

	println("Sum Response:")
	sum(context.Background(), addService, 11111, 22222)

	println("Concat Response:")
	concat(context.Background(), addService, "11111", "22222")
}

func sum(ctx context.Context, svc addsrv.AddService, a, b int) {
	output := svc.Sum(ctx, a, b)
	println(output)
}

func concat(ctx context.Context, svc addsrv.AddService, a, b string) {
	output := svc.Concat(ctx, a, b)
	println(output)
}