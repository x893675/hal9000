package greeter

import (
	"context"
	"fmt"
	"hal9000/proto/greeter"
)

type GreeterServer struct {
}

func NewGreeterServer() *GreeterServer {
	return &GreeterServer{}
}

func (g *GreeterServer) SayHello(ctx context.Context, in *greeter.SayRequest, out *greeter.SayResp) error {
	fmt.Println(in.Msg)
	out.Rsp = "world"
	return nil
}
