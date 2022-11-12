package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config
	config.Add("grpc", map[string]interface{}{
		// Grpc Configuration
		//
		// Configure your server host
		"host": config.Env("GRPC_HOST", ""),

		// Configure your client host and interceptors.
		// Interceptors can be the group name of UnaryClientInterceptorGroups in app/grpc/kernel.go.
		"clients": map[string]any{
			"test": map[string]any{
				"host":         config.Env("GRPC_HOST", ""),
				"interceptors": []string{"test"},
			},
			"timeout": map[string]any{
				"host":         "127.0.0.1:3002",
				"interceptors": []string{"test"},
			},
		},
	})
}
