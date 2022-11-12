package interceptors

import (
	"context"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

func Client(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}

	md["client"] = []string{"goravel-client"}

	if err := invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...); err != nil {
		return err
	}

	return nil
}
