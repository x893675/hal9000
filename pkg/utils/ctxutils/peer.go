package ctxutils

import (
	"context"
	"google.golang.org/grpc/peer"
	"net"
	"strings"
)

func GetClientIP(ctx context.Context) string {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		//logger.Error(ctx, "[getClientIP] invoke FromContext() failed")
		return "unknown"
	}
	if pr.Addr == net.Addr(nil) {
		//logger.Error(ctx, "[getClientIP] peer.Addr is nil")
		return "unknown"
	}

	addSlice := strings.Split(pr.Addr.String(), ":")
	if addSlice[0] == "[" {
		//本机地址
		return "localhost"
	}
	return addSlice[0]
}