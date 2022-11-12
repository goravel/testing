package routes

import (
	"goravel/app/grpc/controllers"
	"goravel/app/grpc/protos"

	"github.com/goravel/framework/facades"
)

func Grpc() {
	protos.RegisterTestServiceServer(facades.Grpc.Server(), &controllers.TestController{})
}
