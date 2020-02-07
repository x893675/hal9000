package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hal9000/internal/addsrv"
	"hal9000/pb"
	"log"
)

func main() {



	conn, err := grpc.Dial(":8891",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*time.Second)))
	//defer cancel()

	client := pb.NewAddClient(conn)
	
	resp, err  := client.Sum(context.Background(), &pb.SumRequest{
		A: 1111,
		B: 2222,
	})

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("client.Search err: deadline")
				return
			}
		}
		log.Println("client.Search err: ", err)
		return
	}

	log.Println("resp is: ", resp.GetV())
	//addService := NewAddService()
	//println("Sum Response:")
	//sum(context.Background(), addService, 11111, 22222)
	//
	//println("Concat Response:")
	//concat(context.Background(), addService, "11111", "22222")
}


func sum(ctx context.Context, svc addsrv.AddService, a, b int) {
	output := svc.Sum(ctx, a, b)
	println(output)
}

func concat(ctx context.Context, svc addsrv.AddService, a, b string) {
	output := svc.Concat(ctx, a, b)
	println(output)
}