package testing

import (
	"testing"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/suite"

	"goravel/bootstrap"
)

type FacadesTestSuite struct {
	suite.Suite
}

func TestFacadesTestSuite(t *testing.T) {
	bootstrap.Boot()

	suite.Run(t, new(FacadesTestSuite))
}

func (s *FacadesTestSuite) SetupTest() {

}

func (s *FacadesTestSuite) TestArtisan() {
	s.NotNil(facades.Artisan)
}

func (s *FacadesTestSuite) TestAuth() {
	s.NotNil(facades.Auth)
}

func (s *FacadesTestSuite) TestCache() {
	s.NotNil(facades.Cache)
}

func (s *FacadesTestSuite) TestConfig() {
	s.NotNil(facades.Config)
}

func (s *FacadesTestSuite) TestOrm() {
	s.NotNil(facades.Orm)
}

func (s *FacadesTestSuite) TestEvent() {
	s.NotNil(facades.Event)
}

func (s *FacadesTestSuite) TestStorage() {
	s.NotNil(facades.Storage)
}

func (s *FacadesTestSuite) TestGrpc() {
	s.NotNil(facades.Grpc)
}

func (s *FacadesTestSuite) TestLog() {
	s.NotNil(facades.Log)
}

func (s *FacadesTestSuite) TestMail() {
	s.NotNil(facades.Mail)
}

func (s *FacadesTestSuite) TestQueue() {
	s.NotNil(facades.Queue)
}

func (s *FacadesTestSuite) TestRoute() {
	s.NotNil(facades.Route)
}

func (s *FacadesTestSuite) TestSchedule() {
	s.NotNil(facades.Schedule)
}

func (s *FacadesTestSuite) TestValidation() {
	s.NotNil(facades.Validation)
}
