package testing

import (
	"testing"

	"goravel/bootstrap"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfigTestSuite(t *testing.T) {
	bootstrap.Boot()

	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) SetupTest() {

}

func (s *ConfigTestSuite) TestConfig() {
	t := s.T()
	assert.Equal(t, "Goravel", facades.Config.GetString("app.name", "laravel"))
	assert.Equal(t, true, facades.Config.GetBool("app.debug", false))
	assert.Equal(t, 587, facades.Config.GetInt("mail.port", 123))
	assert.Equal(t, "Goravel", facades.Config.Env("APP_NAME", "laravel").(string))
	assert.Equal(t, "127.0.0.1:3001", facades.Config.Env("GRPC_HOST", "Goravel").(string))
}
