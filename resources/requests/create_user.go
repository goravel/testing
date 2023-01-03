package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateUser struct {
	Name string `form:"name" json:"name"`
}

func (r *CreateUser) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateUser) Rules() map[string]string {
	return map[string]string{
		"name": "required",
	}
}

func (r *CreateUser) Messages() map[string]string {
	return map[string]string{}
}

func (r *CreateUser) Attributes() map[string]string {
	return map[string]string{}
}

func (r *CreateUser) PrepareForValidation(data validation.Data) {
	if name, exist := data.Get("name"); exist {
		_ = data.Set("name", name.(string)+"1")
	}
}
