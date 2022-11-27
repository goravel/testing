package testing

import (
	"context"
	"fmt"
	"github.com/goravel/framework/facades"
	"goravel/bootstrap"
	"testing"

	"github.com/goravel/framework/contracts/auth/access"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GateTestSuite struct {
	suite.Suite
}

func TestGateTestSuite(t *testing.T) {
	bootstrap.Boot()

	suite.Run(t, new(GateTestSuite))
}

func (s *GateTestSuite) SetupTest() {

}

func (s *GateTestSuite) TestWithContext() {
	ctx := context.WithValue(context.Background(), "hello", "goravel")

	assert.Equal(s.T(), access.NewDenyResponse("goravel"), facades.Gate.WithContext(ctx).Inspect("context", map[string]any{
		"user": "2",
	}))
}

func (s *GateTestSuite) TestAllows() {
	assert.True(s.T(), facades.Gate.Allows("create", map[string]any{
		"user": "1",
	}))
	assert.False(s.T(), facades.Gate.Allows("create", map[string]any{
		"user": "2",
	}))
	assert.False(s.T(), facades.Gate.Allows("update", map[string]any{
		"user": "1",
	}))
}

func (s *GateTestSuite) TestDenies() {
	assert.False(s.T(), facades.Gate.Denies("create", map[string]any{
		"user": "1",
	}))
	assert.True(s.T(), facades.Gate.Denies("create", map[string]any{
		"user": "2",
	}))
	assert.True(s.T(), facades.Gate.Denies("update", map[string]any{
		"user": "1",
	}))
}

func (s *GateTestSuite) TestInspect() {
	assert.Equal(s.T(), access.NewAllowResponse(), facades.Gate.Inspect("create", map[string]any{
		"user": "1",
	}))
	assert.True(s.T(), facades.Gate.Inspect("create", map[string]any{
		"user": "1",
	}).Allowed())
	assert.Equal(s.T(), access.NewDenyResponse("create error"), facades.Gate.Inspect("create", map[string]any{
		"user": "2",
	}))
	assert.Equal(s.T(), "create error", facades.Gate.Inspect("create", map[string]any{
		"user": "2",
	}).Message())
	assert.Equal(s.T(), access.NewDenyResponse(fmt.Sprintf("ability doesn't exist: %s", "deletes")), facades.Gate.Inspect("deletes", map[string]any{
		"user": "1",
	}))
}

func (s *GateTestSuite) TestAny() {
	assert.True(s.T(), facades.Gate.Any([]string{"create", "update"}, map[string]any{
		"user": "1",
	}))
	assert.True(s.T(), facades.Gate.Any([]string{"create", "update"}, map[string]any{
		"user": "2",
	}))
	assert.False(s.T(), facades.Gate.Any([]string{"create", "update"}, map[string]any{
		"user": "4",
	}))
}

func (s *GateTestSuite) TestNone() {
	assert.False(s.T(), facades.Gate.None([]string{"create", "update"}, map[string]any{
		"user": "1",
	}))
	assert.False(s.T(), facades.Gate.None([]string{"create", "update"}, map[string]any{
		"user": "2",
	}))
	assert.True(s.T(), facades.Gate.None([]string{"create", "update"}, map[string]any{
		"user": "4",
	}))
}

func (s *GateTestSuite) TestBefore() {
	assert.True(s.T(), facades.Gate.Allows("create", map[string]any{
		"user": "3",
	}))
	assert.False(s.T(), facades.Gate.Allows("create", map[string]any{
		"user": "4",
	}))
}

func (s *GateTestSuite) TestAfter() {
	assert.True(s.T(), facades.Gate.Allows("delete", map[string]any{
		"user": "1",
	}))
	assert.True(s.T(), facades.Gate.Allows("delete", map[string]any{
		"user": "3",
	}))
}
