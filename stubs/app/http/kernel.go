package http

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/http/middleware"
)

type Kernel struct {
}

func (kernel *Kernel) Middleware() []http.Middleware {
	return []http.Middleware{
		middleware.Cors(),
		func(ctx http.Context) {
			ctx.WithValue("global", "goravel")
			ctx.Request().Next()
		},
	}
}
