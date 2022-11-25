package routes

import (
	"encoding/json"
	nethttp "net/http"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Web() {
	// ------------------
	// Test Route
	// ------------------
	facades.Route.Prefix("group1").Middleware(TestContextMiddleware()).Group(func(route1 route.Route) {
		route1.Prefix("group2").Middleware(TestContextMiddleware1()).Group(func(route2 route.Route) {
			route2.Get("/middleware/{id}", func(ctx http.Context) {
				ctx.Response().Success().Json(http.Json{
					"id":   ctx.Request().Input("id"),
					"ctx":  ctx.Value("ctx").(string),
					"ctx1": ctx.Value("ctx1").(string),
				})
			})
		})
		route1.Middleware(TestContextMiddleware2()).Get("/middleware/{id}", func(ctx http.Context) {
			ctx.Response().Success().Json(http.Json{
				"id":   ctx.Request().Input("id"),
				"ctx":  ctx.Value("ctx").(string),
				"ctx2": ctx.Value("ctx2").(string),
			})
		})
	})

	facades.Route.Get("/input/{id}", func(ctx http.Context) {
		ctx.Response().Json(nethttp.StatusOK, http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Post("/input/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Put("/input/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Delete("/input/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Options("/input/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Patch("/input/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Any("/any/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Static("static", "./resources")
	facades.Route.StaticFile("static-file", "./resources/logo.png")
	facades.Route.StaticFS("static-fs", nethttp.Dir("./public"))

	facades.Route.Middleware(TestAbortMiddleware()).Get("/middleware/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Middleware(TestContextMiddleware(), TestContextMiddleware1()).Get("/middlewares/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id":   ctx.Request().Input("id"),
			"ctx":  ctx.Value("ctx"),
			"ctx1": ctx.Value("ctx1"),
		})
	})

	facades.Route.Prefix("prefix1").Prefix("prefix2").Get("input/{id}", func(ctx http.Context) {
		ctx.Response().Success().Json(http.Json{
			"id": ctx.Request().Input("id"),
		})
	})

	facades.Route.Get("/global-middleware", func(ctx http.Context) {
		ctx.Response().Json(nethttp.StatusOK, http.Json{
			"global": ctx.Value("global"),
		})
	})

	// ------------------
	// Test Request
	// ------------------
	facades.Route.Prefix("request").Group(func(route route.Route) {
		route.Get("/get/{id}", func(ctx http.Context) {
			ctx.Response().Success().Json(http.Json{
				"id":       ctx.Request().Input("id"),
				"name":     ctx.Request().Query("name", "Hello"),
				"header":   ctx.Request().Header("Hello", "World"),
				"method":   ctx.Request().Method(),
				"path":     ctx.Request().Path(),
				"url":      ctx.Request().Url(),
				"full_url": ctx.Request().FullUrl(),
				"ip":       ctx.Request().Ip(),
			})
		})
		route.Get("/headers", func(ctx http.Context) {
			str, _ := json.Marshal(ctx.Request().Headers())
			ctx.Response().Success().String(string(str))
		})
		route.Post("/post", func(ctx http.Context) {
			ctx.Response().Success().Json(http.Json{
				"name": ctx.Request().Form("name", "Hello"),
			})
		})
		route.Post("/bind", func(ctx http.Context) {
			type Test struct {
				Name string
			}
			var test Test
			_ = ctx.Request().Bind(&test)
			ctx.Response().Success().Json(http.Json{
				"name": test.Name,
			})
		})
		route.Post("/file", func(ctx http.Context) {
			file, err := ctx.Request().File("file")
			if err != nil {
				ctx.Response().Success().String("get file error")
				return
			}
			filePath, err := file.Store("test")
			if err != nil {
				ctx.Response().Success().String("store file error: " + err.Error())
				return
			}

			extension, err := file.Extension()
			if err != nil {
				ctx.Response().Success().String("get file extension error: " + err.Error())
				return
			}

			ctx.Response().Success().Json(http.Json{
				"exist":              facades.Storage.Exists(filePath),
				"hash_name_length":   len(file.HashName()),
				"hash_name_length1":  len(file.HashName("test")),
				"file_path_length":   len(filePath),
				"extension":          extension,
				"original_name":      file.GetClientOriginalName(),
				"original_extension": file.GetClientOriginalExtension(),
			})
		})
	})

	// ------------------
	// Test Response
	// ------------------
	facades.Route.Prefix("response").Group(func(route route.Route) {
		route.Get("/json", func(ctx http.Context) {
			ctx.Response().Json(nethttp.StatusOK, http.Json{
				"id": "1",
			})
		})
		route.Get("/string", func(ctx http.Context) {
			ctx.Response().String(nethttp.StatusCreated, "Goravel")
		})
		route.Get("/success/json", func(ctx http.Context) {
			ctx.Response().Success().Json(http.Json{
				"id": "1",
			})
		})
		route.Get("/success/string", func(ctx http.Context) {
			ctx.Response().Success().String("Goravel")
		})
		route.Get("/file", func(ctx http.Context) {
			ctx.Response().File("./resources/logo.png")
		})
		route.Get("/download", func(ctx http.Context) {
			ctx.Response().Download("./resources/logo.png", "1.png")
		})
		route.Get("/header", func(ctx http.Context) {
			ctx.Response().Header("Hello", "goravel").String(nethttp.StatusOK, "Goravel")
		})
	})
}

func TestAbortMiddleware() http.Middleware {
	return func(ctx http.Context) {
		ctx.Request().AbortWithStatus(nethttp.StatusNonAuthoritativeInfo)
		return
	}
}

func TestContextMiddleware() http.Middleware {
	return func(ctx http.Context) {
		ctx.WithValue("ctx", "Goravel")

		ctx.Request().Next()
	}
}

func TestContextMiddleware1() http.Middleware {
	return func(ctx http.Context) {
		ctx.WithValue("ctx1", "Hello")

		ctx.Request().Next()
	}
}

func TestContextMiddleware2() http.Middleware {
	return func(ctx http.Context) {
		ctx.WithValue("ctx2", "World")

		ctx.Request().Next()
	}
}
