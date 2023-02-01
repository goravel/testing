package testing

import (
	"testing"

	"goravel/bootstrap"

	"github.com/goravel/framework/support/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArtisanTestSuite struct {
	suite.Suite
}

func TestArtisanTestSuite(t *testing.T) {
	bootstrap.Boot()

	suite.Run(t, new(ArtisanTestSuite))
}

func (s *ArtisanTestSuite) SetupTest() {

}

func (s *ArtisanTestSuite) TestCacheClear() {
	Equal(s.T(), "cache:clear", "Application cache cleared")
}

func (s *ArtisanTestSuite) TestHelp() {
	NotEmpty(s.T(), "help migrate")
}

func (s *ArtisanTestSuite) TestList() {
	NotEmpty(s.T(), "list")
}

func (s *ArtisanTestSuite) TestJwtSecret() {
	t := s.T()
	Equal(t, "jwt:secret", "Jwt Secret set successfully")
	Equal(t, "jwt:secret", "Exist jwt secret")
}

func (s *ArtisanTestSuite) TestKeyGenerate() {
	t := s.T()
	Equal(t, "key:generate", "Application key set successfully")
	Equal(t, "key:generate", "Exist application key")
}

func (s *ArtisanTestSuite) TestMakeCommand() {
	t := s.T()
	Equal(t, "make:command SendEmails", "Console command created successfully")
	assert.True(t, file.Exists("./app/console/commands/send_emails.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *ArtisanTestSuite) TestMakeEvent() {
	t := s.T()
	Equal(t, "make:event OrderShipped", "Event created successfully")
	assert.True(t, file.Exists("./app/events/order_shipped.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *ArtisanTestSuite) TestMakeJob() {
	t := s.T()
	Equal(t, "make:job TestJob", "Job created successfully")
	assert.True(t, file.Exists("./app/jobs/test_job.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *ArtisanTestSuite) TestMakeListener() {
	t := s.T()
	Equal(t, "make:listener SendShipmentNotification", "Listener created successfully")
	assert.True(t, file.Exists("./app/listeners/send_shipment_notification.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *ArtisanTestSuite) TestMakeMigration() {
	t := s.T()
	Equal(t, "make:migration create_users_table", "Created Migration: create_users_table")
	assert.True(t, file.Exists("./database/migrations"))
	assert.True(t, file.Remove("./database"))
}

func (s *ArtisanTestSuite) TestMakePolicy() {
	t := s.T()
	Equal(t, "make:policy UserPolicy", "Policy created successfully")
	assert.True(t, file.Exists("./app/policies/user_policy.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *ArtisanTestSuite) TestMakeRequest() {
	t := s.T()
	Equal(t, "make:request UserRequest", "Request created successfully")
	assert.True(t, file.Exists("./app/http/requests/user_request.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *ArtisanTestSuite) TestMakeRule() {
	t := s.T()
	Equal(t, "make:rule UserRule", "Rule created successfully")
	assert.True(t, file.Exists("./app/rules/user_rule.go"))
	assert.True(t, file.Remove("./app"))
}
