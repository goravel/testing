package controllers

import (
	"context"
	"errors"
	"fmt"
	"goravel/app/grpc/protos"
	"net/http"
)

type TestController struct {
}

func (r *TestController) Get(ctx context.Context, req *protos.TestRequest) (*protos.TestResponse, error) {
	if req.GetName() == "success" {
		return &protos.TestResponse{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Goravel: server: %s, client: %s", ctx.Value("server"), ctx.Value("client")),
		}, nil
	} else {
		return nil, errors.New("error")
	}
}
