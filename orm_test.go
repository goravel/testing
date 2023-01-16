package testing

import (
	"testing"

	"goravel/app/models"
	"goravel/bootstrap"
	testingmodels "goravel/testing/resources/models"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrmTestSuite struct {
	suite.Suite
	start bool
}

func TestOrmTestSuite(t *testing.T) {
	bootstrap.Boot()
	suite.Run(t, new(OrmTestSuite))
}

func (s *OrmTestSuite) SetupTest() {
	if !s.start {
		migrate(s.T())
		s.start = true
	}
}

func (s *OrmTestSuite) TestMakeMigration() {
	t := s.T()
	Equal(t, "make:migration create_users_table", "Created Migration: create_users_table")
	assert.True(t, file.Exists("./database/migrations"))
	assert.True(t, file.Remove("./database"))
}

func migrate(t *testing.T) {
	clearTables(t)

	outStr, errStr, err := RunCommand("cp -R stubs/database database")
	assert.Empty(t, outStr)
	assert.Empty(t, errStr)
	assert.NoError(t, err)

	Equal(t, "migrate", "Migration success")
	user := models.User{Name: "user"}
	assert.Nil(t, facades.Orm.Query().Create(&user))
	assert.Nil(t, facades.Orm.Query().Create(&testingmodels.UserAddress{Name: "address", UserId: user.ID}))

	clearData(t)
	file.Remove("./database")
}

func clearTables(t *testing.T) {
	facades.Orm.Query().Exec("DROP TABLE users")
	facades.Orm.Query().Exec("DROP TABLE user_addresses;")
	facades.Orm.Query().Exec("DROP TABLE migrations;")
}

func clearData(t *testing.T) {
	assert.Nil(t, facades.Orm.Query().Exec("TRUNCATE table users"))
	assert.Nil(t, facades.Orm.Query().Exec("TRUNCATE table user_addresses"))
}
