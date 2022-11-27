package providers

import (
	"context"
	"goravel/testing/resources/policies"

	"github.com/goravel/framework/contracts/auth/access"
	"github.com/goravel/framework/facades"
)

type AuthServiceProvider struct {
}

func (receiver *AuthServiceProvider) Register() {
}

func (receiver *AuthServiceProvider) Boot() {
	facades.Gate.Before(func(ctx context.Context, ability string, arguments map[string]any) *access.Response {
		user := arguments["user"].(string)
		if user == "3" {
			return access.NewAllowResponse()
		}

		return nil
	})

	facades.Gate.After(func(ctx context.Context, ability string, arguments map[string]any, result *access.Response) *access.Response {
		user := arguments["user"].(string)
		if user == "4" {
			return access.NewAllowResponse()
		}

		return nil
	})

	facades.Gate.Define("context", func(ctx context.Context, arguments map[string]any) *access.Response {
		user := arguments["user"].(string)
		if user == "1" {
			return access.NewAllowResponse()
		} else {
			return access.NewDenyResponse(ctx.Value("hello").(string))
		}
	})

	facades.Gate.Define("create", func(ctx context.Context, arguments map[string]any) *access.Response {
		user := arguments["user"].(string)
		if user == "1" {
			return access.NewAllowResponse()
		} else {
			return access.NewDenyResponse("create error")
		}
	})

	facades.Gate.Define("delete", func(ctx context.Context, arguments map[string]any) *access.Response {
		user := arguments["user"].(string)
		if user == "3" {
			return nil
		} else {
			return access.NewAllowResponse()
		}
	})

	facades.Gate.Define("update", policies.NewUserPolicy().Update)
}
