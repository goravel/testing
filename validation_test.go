package testing

import (
	"testing"

	"goravel/bootstrap"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/suite"
)

type ValidationTestSuite struct {
	suite.Suite
}

func TestValidationTestSuite(t *testing.T) {
	bootstrap.Boot()

	suite.Run(t, new(ValidationTestSuite))
}

func (s *ValidationTestSuite) SetupTest() {
}

func (s *ValidationTestSuite) TestValidatorSuccess() {
	bootstrap.Boot()
	validator, err := facades.Validation.Make(
		map[string]any{"a": "aa", "b": 1},
		map[string]string{"a": "required", "b": "required"},
	)
	s.Nil(err)
	type Data struct {
		A string
		B int
	}
	var data Data
	err = validator.Bind(&data)
	s.Nil(err)
	s.Equal("aa", data.A)
	s.Equal(1, data.B)
}

func (s *ValidationTestSuite) TestValidatorError() {
	bootstrap.Boot()
	validator, err := facades.Validation.Make(
		map[string]any{"a": "aa"},
		map[string]string{"a": "required", "b": "required"},
	)
	s.Nil(err)
	type Data struct {
		A string
		B int
	}
	var data Data
	err = validator.Bind(&data)
	s.Nil(err)
	s.Equal("aa", data.A)
	s.Equal(0, data.B)
	s.Equal("b is required to not be empty", validator.Errors().One())
	s.True(validator.Fails())
}
