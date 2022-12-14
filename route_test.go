package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"goravel/bootstrap"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	bootstrap.Boot()
	tests := []struct {
		name       string
		setup      func(req *http.Request)
		method     string
		url        string
		expectCode int
		expectBody string
	}{
		{
			name:       "Get",
			method:     "GET",
			url:        "/input/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Post",
			method:     "POST",
			url:        "/input/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Put",
			method:     "PUT",
			url:        "/input/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Delete",
			method:     "DELETE",
			url:        "/input/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Options",
			method:     "OPTIONS",
			url:        "/input/1",
			expectCode: http.StatusOK,
		},
		{
			name:       "Patch",
			method:     "PATCH",
			url:        "/input/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Any Get",
			method:     "GET",
			url:        "/any/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Any Post",
			method:     "POST",
			url:        "/any/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Any Put",
			method:     "PUT",
			url:        "/any/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Any Delete",
			method:     "DELETE",
			url:        "/any/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name: "Any Options",
			setup: func(req *http.Request) {
				req.Header.Set("Origin", "http://127.0.0.1")
				req.Header.Set("Access-Control-Request-Method", "GET")
			},
			method:     "OPTIONS",
			url:        "/any/1",
			expectCode: http.StatusNoContent,
		},
		{
			name:       "Any Patch",
			method:     "PATCH",
			url:        "/any/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Static",
			method:     "GET",
			url:        "/static/logo.png",
			expectCode: http.StatusOK,
		},
		{
			name:       "StaticFile",
			method:     "GET",
			url:        "/static-file",
			expectCode: http.StatusOK,
		},
		{
			name:       "StaticFS",
			method:     "GET",
			url:        "/static-fs",
			expectCode: http.StatusMovedPermanently,
		},
		{
			name:       "Abort Middleware",
			method:     "GET",
			url:        "/middleware/1",
			expectCode: http.StatusNonAuthoritativeInfo,
		},
		{
			name:       "Multiple Middleware",
			method:     "GET",
			url:        "/middlewares/1",
			expectCode: http.StatusOK,
			expectBody: "{\"ctx\":\"Goravel\",\"ctx1\":\"Hello\",\"id\":\"1\"}",
		},
		{
			name:       "Multiple Prefix",
			method:     "GET",
			url:        "/prefix1/prefix2/input/1",
			expectCode: http.StatusOK,
			expectBody: "{\"id\":\"1\"}",
		},
		{
			name:       "Multiple Prefix Group Middleware",
			method:     "GET",
			url:        "/group1/group2/middleware/1",
			expectCode: http.StatusOK,
			expectBody: "{\"ctx\":\"Goravel\",\"ctx1\":\"Hello\",\"id\":\"1\"}",
		},
		{
			name:       "Multiple Group Middleware",
			method:     "GET",
			url:        "/group1/middleware/1",
			expectCode: http.StatusOK,
			expectBody: "{\"ctx\":\"Goravel\",\"ctx2\":\"World\",\"id\":\"1\"}",
		},
		{
			name:       "Global Middleware",
			method:     "GET",
			url:        "/global-middleware",
			expectCode: http.StatusOK,
			expectBody: "{\"global\":\"goravel\"}",
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(test.method, test.url, nil)
		if test.setup != nil {
			test.setup(req)
		}
		facades.Route.ServeHTTP(w, req)

		if test.expectBody != "" {
			assert.Equal(t, test.expectBody, w.Body.String(), test.name)
		}
		assert.Equal(t, test.expectCode, w.Code, test.name)
	}
}
