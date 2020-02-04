package main

import (
	"context"
	"hal9000/internal/addsrv"
)

func main() {

	addService := NewAddService()
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