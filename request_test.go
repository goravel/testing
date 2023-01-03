package testing

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"goravel/bootstrap"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/file"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	bootstrap.Boot()

	var (
		req *http.Request
	)

	tests := []struct {
		name       string
		method     string
		url        string
		setup      func(method, url string) error
		expectCode int
		expectBody string
	}{
		{
			name:   "Methods",
			method: "GET",
			url:    "/request/get/1?name=Goravel",
			setup: func(method, url string) error {
				var err error
				req, err = http.NewRequest(method, url, nil)
				if err != nil {
					return err
				}
				req.Header.Set("Hello", "goravel")

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"full_url\":\"\",\"header\":\"goravel\",\"id\":\"1\",\"ip\":\"\",\"method\":\"GET\",\"name\":\"Goravel\",\"path\":\"/request/get/1\",\"url\":\"\"}",
		},
		{
			name:   "Headers",
			method: "GET",
			url:    "/request/headers",
			setup: func(method, url string) error {
				var err error
				req, err = http.NewRequest(method, url, nil)
				if err != nil {
					return err
				}
				req.Header.Set("Hello", "Goravel")

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"Hello\":[\"Goravel\"]}",
		},
		{
			name:   "Form",
			method: "POST",
			url:    "/request/post",
			setup: func(method, url string) error {
				payload := &bytes.Buffer{}
				writer := multipart.NewWriter(payload)
				if err := writer.WriteField("name", "Goravel"); err != nil {
					return err
				}
				if err := writer.Close(); err != nil {
					return err
				}

				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", writer.FormDataContentType())

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"name\":\"Goravel\"}",
		},
		{
			name:   "Bind",
			method: "POST",
			url:    "/request/bind",
			setup: func(method, url string) error {
				payload := strings.NewReader(`{
					"Name": "Goravel"
				}`)
				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", "application/json")

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"name\":\"Goravel\"}",
		},
		{
			name:   "File",
			method: "POST",
			url:    "/request/file",
			setup: func(method, url string) error {
				payload := &bytes.Buffer{}
				writer := multipart.NewWriter(payload)
				logo, errFile1 := os.Open("./resources/logo.png")
				defer logo.Close()
				part1, errFile1 := writer.CreateFormFile("file", filepath.Base("./resources/logo.png"))
				_, errFile1 = io.Copy(part1, logo)
				if errFile1 != nil {
					return errFile1
				}
				err := writer.Close()
				if err != nil {
					return err
				}

				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", writer.FormDataContentType())

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"exist\":true,\"extension\":\"png\",\"file_path_length\":49,\"hash_name_length\":44,\"hash_name_length1\":49,\"original_extension\":\"png\",\"original_name\":\"logo.png\"}",
		},
		{
			name:   "GET with validator and validate pass",
			method: "GET",
			url:    "/request/validator/validate/success?name=Goravel",
			setup: func(method, url string) error {
				var err error
				req, err = http.NewRequest(method, url, nil)
				if err != nil {
					return err
				}

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"name\":\"Goravel\"}",
		},
		{
			name:   "GET with validator but validate fail",
			method: "GET",
			url:    "/request/validator/validate/fail?name=Goravel",
			setup: func(method, url string) error {
				var err error
				req, err = http.NewRequest(method, url, nil)
				if err != nil {
					return err
				}

				return nil
			},
			expectCode: http.StatusBadRequest,
			expectBody: "Validate fail: map[name1:map[required:name1 is required to not be empty]]",
		},
		{
			name:   "GET with validator and validate request pass",
			method: "GET",
			url:    "/request/validator/validate-request/success?name=Goravel",
			setup: func(method, url string) error {
				var err error
				req, err = http.NewRequest(method, url, nil)
				if err != nil {
					return err
				}

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"name\":\"Goravel1\"}",
		},
		{
			name:   "GET with validator but validate request fail",
			method: "GET",
			url:    "/request/validator/validate-request/fail?name1=Goravel",
			setup: func(method, url string) error {
				var err error
				req, err = http.NewRequest(method, url, nil)
				if err != nil {
					return err
				}

				return nil
			},
			expectCode: http.StatusBadRequest,
			expectBody: "Validate fail: map[name:map[required:name is required to not be empty]]",
		},
		{
			name:   "POST with validator and validate pass",
			method: "POST",
			url:    "/request/validator/validate/success",
			setup: func(method, url string) error {
				payload := strings.NewReader(`{
					"name": "Goravel"
				}`)
				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", "application/json")

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"name\":\"Goravel\"}",
		},
		{
			name:   "POST with validator and validate fail",
			method: "POST",
			url:    "/request/validator/validate/fail",
			setup: func(method, url string) error {
				payload := strings.NewReader(`{
					"name": "Goravel"
				}`)
				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", "application/json")

				return nil
			},
			expectCode: http.StatusBadRequest,
			expectBody: "Validate fail: map[name1:map[required:name1 is required to not be empty]]",
		},
		{
			name:   "POST with validator and validate request pass",
			method: "POST",
			url:    "/request/validator/validate-request/success",
			setup: func(method, url string) error {
				payload := strings.NewReader(`{
					"name": "Goravel"
				}`)
				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", "application/json")

				return nil
			},
			expectCode: http.StatusOK,
			expectBody: "{\"name\":\"Goravel1\"}",
		},
		{
			name:   "POST with validator and validate request fail",
			method: "POST",
			url:    "/request/validator/validate-request/fail",
			setup: func(method, url string) error {
				payload := strings.NewReader(`{
					"name1": "Goravel"
				}`)
				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", "application/json")

				return nil
			},
			expectCode: http.StatusBadRequest,
			expectBody: "Validate fail: map[name:map[required:name is required to not be empty]]",
		},
		{
			name:   "POST with validator and validate request unauthorize",
			method: "POST",
			url:    "/request/validator/validate-request/unauthorize",
			setup: func(method, url string) error {
				payload := strings.NewReader(`{
					"name": "Goravel"
				}`)
				req, _ = http.NewRequest(method, url, payload)
				req.Header.Set("Content-Type", "application/json")

				return nil
			},
			expectCode: http.StatusBadRequest,
			expectBody: "Validate error: error",
		},
	}

	for _, test := range tests {
		err := test.setup(test.method, test.url)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		facades.Route.ServeHTTP(w, req)

		if test.expectBody != "" {
			assert.Equal(t, test.expectBody, w.Body.String(), test.name)
		}
		assert.Equal(t, test.expectCode, w.Code, test.name)

		file.Remove("./storage")
	}
}
