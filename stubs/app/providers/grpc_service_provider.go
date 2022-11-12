package providers

import (
	"goravel/app/grpc"
	"goravel/routes"

	"github.com/goravel/framework/facades"
)

type GrpcServiceProvider struct {
}

func (receiver *GrpcServiceProvider) Register() {
	kernel := grpc.Kernel{}
	facades.Grpc.UnaryServerInterceptors(kernel.UnaryServerInterceptors())
	facades.Grpc.UnaryClientInterceptorGroups(kernel.UnaryClientInterceptorGroups())
}

func (receiver *GrpcServiceProvider) Boot() {
	routes.Grpc()
}
