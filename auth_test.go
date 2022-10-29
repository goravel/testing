package testing

import (
	"testing"

	"goravel/app/models"
	"goravel/bootstrap"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"
	"github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
}

func TestAuthTestSuite(t *testing.T) {
	facades.Config.Add("jwt", map[string]interface{}{
		"secret":      "Goravel",
		"ttl":         60,
		"refresh_ttl": 20160,
	})

	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) SetupTest() {
	bootstrap.Boot()
}

func (s *AuthTestSuite) TestJwtSecret() {
	t := s.T()
	Equal(t, "jwt:secret", "Jwt Secret set successfully")
	Equal(t, "jwt:secret", "Exist jwt secret")
}

func (s *AuthTestSuite) TestLogin_NoPrimaryKey() {
	type User struct {
		ID   uint
		Name string
	}

	ctx := http.Background()
	var user User
	user.ID = 1
	user.Name = "Goravel"
	token, err := facades.Auth.Login(ctx, &user)
	assert.Empty(s.T(), token)
	assert.ErrorIs(s.T(), err, auth.ErrorNoPrimaryKeyField)
}

func (s *AuthTestSuite) TestLogin() {
	ctx := http.Background()
	var user models.User
	user.ID = 1
	user.Name = "Goravel"
	token, err := facades.Auth.Login(ctx, &user)
	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestLoginUsingID() {
	ctx := http.Background()
	token, err := facades.Auth.LoginUsingID(ctx, 1)
	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestParse_InvalidToken() {
	ctx := http.Background()
	err := facades.Auth.Parse(ctx, "1")
	assert.NotNil(s.T(), err)
}

func (s *AuthTestSuite) TestParse() {
	ctx := http.Background()
	token, err := facades.Auth.LoginUsingID(ctx, 1)
	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)

	err = facades.Auth.Parse(ctx, token)
	assert.Nil(s.T(), err)

	err = facades.Auth.Parse(ctx, "Bearer "+token)
	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestUser_NoParse() {
	ctx := http.Background()
	var user models.User
	err := facades.Auth.User(ctx, &user)
	assert.NotNil(s.T(), err)
}

func (s *AuthTestSuite) TestUser() {
	ctx := http.Background()
	token, err := facades.Auth.LoginUsingID(ctx, 1)
	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)

	var user models.User
	mockOrm, mockDB, _ := mock.Orm()
	mockOrm.On("Query").Return(mockDB).Once()
	mockDB.On("Find", &user, "1").Return(nil).Once()

	err = facades.Auth.User(ctx, &user)
	assert.Nil(s.T(), err)

	err = facades.Auth.Parse(ctx, token)
	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestRefresh_NoParse() {
	ctx := http.Background()
	token, err := facades.Auth.Refresh(ctx)
	assert.Empty(s.T(), token)
	assert.NotNil(s.T(), err)
}

func (s *AuthTestSuite) TestRefresh() {
	ctx := http.Background()
	token, err := facades.Auth.LoginUsingID(ctx, 1)
	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)

	token, err = facades.Auth.Refresh(ctx)
	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)
}

func (s *AuthTestSuite) TestLogout() {
	ctx := http.Background()
	token, err := facades.Auth.Guard("admin").LoginUsingID(ctx, 1)
	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)

	var user models.User
	mockOrm, mockDB, _ := mock.Orm()
	mockOrm.On("Query").Return(mockDB).Once()
	mockDB.On("Find", &user, "1").Return(nil).Once()

	err = facades.Auth.Guard("admin").User(ctx, &user)
	assert.Nil(s.T(), err)

	err = facades.Auth.Guard("admin").Logout(ctx)
	assert.Nil(s.T(), err)

	err = facades.Auth.Guard("admin").User(ctx, &user)
	assert.NotNil(s.T(), err)

	err = facades.Auth.Guard("admin").Parse(ctx, token)
	assert.ErrorIs(s.T(), err, auth.ErrorTokenDisabled)
}
