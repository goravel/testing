package policies

import (
	"context"

	"github.com/goravel/framework/contracts/auth/access"
)

type UserPolicy struct {
}

func NewUserPolicy() *UserPolicy {
	return &UserPolicy{}
}

func (r *UserPolicy) Update(ctx context.Context, arguments map[string]any) *access.Response {
	user := arguments["user"].(string)
	if user == "2" {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse(" update error")
	}
}
