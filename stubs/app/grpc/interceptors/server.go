package interceptors

import (
	"context"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

func Server(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	ctx = context.WithValue(ctx, "server", "goravel-server")
	if len(md["client"]) > 0 {
		ctx = context.WithValue(ctx, "client", md["client"][0])
	}

	return handler(ctx, req)
}
