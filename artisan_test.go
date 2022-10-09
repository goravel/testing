package testing

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"goravel/bootstrap"
	"io/ioutil"

	"strings"
	"testing"
	"time"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/file"
	"github.com/stretchr/testify/assert"
)

type ArtisanTestSuite struct {
	suite.Suite
}

func TestArtisanTestSuite(t *testing.T) {
	file.Remove("./storage")

	bootstrap.Boot()

	suite.Run(t, new(ArtisanTestSuite))
}

func (s *ArtisanTestSuite) SetupTest() {

}

func (s *ArtisanTestSuite) TestKeyGenerate() {
	t := s.T()
	Equal(t, "key:generate", "Application key set successfully")
	Equal(t, "key:generate", "Exist application key")
}

func (s *ArtisanTestSuite) TestList() {
	t := s.T()
	NotEmpty(t, "list")
}

func (s *ArtisanTestSuite) TestHelp() {
	t := s.T()
	NotEmpty(t, "help migrate")
}

func (s *ArtisanTestSuite) TestMakeCommand() {
	t := s.T()
	Equal(t, "make:command SendEmails", "Console command created successfully")
	assert.True(t, file.Exist("./app/console/commands/send_emails.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *ArtisanTestSuite) TestCommand() {
	t := s.T()
	facades.Artisan.Call("test --name Goravel argument0 argument1")
	facades.Artisan.Call("test -n Goravel1 --age 20 argument2 argument3")

	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	assert.True(t, file.Exist(log))
	data, err := ioutil.ReadFile(log)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(string(data), "Run test command success, argument_0: argument0, argument_1: argument1, option_name: Goravel, option_age: 18, arguments: argument0,argument1"))
	assert.True(t, strings.Contains(string(data), "Run test command success, argument_0: argument2, argument_1: argument3, option_name: Goravel1, option_age: 20, arguments: argument2,argument3"))
}
