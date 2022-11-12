package grpc

import (
	"goravel/app/grpc/interceptors"

	"google.golang.org/grpc"
)

type Kernel struct {
}

// The application's global GRPC middleware stack.
// These middleware are run during every request to your application.
func (kernel *Kernel) UnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		interceptors.Server,
	}
}

func (kernel *Kernel) UnaryClientInterceptorGroups() map[string][]grpc.UnaryClientInterceptor {
	return map[string][]grpc.UnaryClientInterceptor{
		"test": {
			interceptors.Client,
		},
	}
}
