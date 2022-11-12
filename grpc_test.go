package testing

import (
	"context"
	"net/http"
	"testing"

	"goravel/app/grpc/protos"
	"goravel/bootstrap"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GrpcTestSuite struct {
	suite.Suite
}

func TestGrpcTestSuite(t *testing.T) {
	bootstrap.Boot()

	suite.Run(t, new(GrpcTestSuite))
}

func (s *GrpcTestSuite) SetupTest() {

}

func (s *GrpcTestSuite) TestSuccess() {
	go func() {
		if err := facades.Grpc.Run(facades.Config.GetString("grpc.host")); err != nil {
			facades.Log.Error("Run grpc error")
		}
	}()

	client, err := facades.Grpc.Client(context.Background(), "test")
	assert.Nil(s.T(), err)
	testServiceClient := protos.NewTestServiceClient(client)
	res, err := testServiceClient.Get(context.Background(), &protos.TestRequest{
		Name: "success",
	})

	assert.Equal(s.T(), &protos.TestResponse{Code: http.StatusOK, Message: "Goravel: server: goravel-server, client: goravel-client"}, res)
	assert.Nil(s.T(), err)
}

func (s *GrpcTestSuite) TestError() {
	go func() {
		if err := facades.Grpc.Run(facades.Config.GetString("grpc.host")); err != nil {
			facades.Log.Error("Run grpc error")
		}
	}()

	client, err := facades.Grpc.Client(context.Background(), "test")
	assert.Nil(s.T(), err)
	testServiceClient := protos.NewTestServiceClient(client)
	res, err := testServiceClient.Get(context.Background(), &protos.TestRequest{
		Name: "error",
	})

	assert.Nil(s.T(), res)
	assert.EqualError(s.T(), err, "rpc error: code = Unknown desc = error")
}

func (s *GrpcTestSuite) TestErrorWhenServerDoesntRun() {
	client, err := facades.Grpc.Client(context.Background(), "timeout")
	assert.Nil(s.T(), client)
	assert.ErrorIs(s.T(), err, context.DeadlineExceeded)
}
