package requests

import (
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type Unauthorize struct {
	Name string `form:"name" json:"name"`
}

func (r *Unauthorize) Authorize(ctx http.Context) error {
	return errors.New("error")
}

func (r *Unauthorize) Rules() map[string]string {
	return map[string]string{
		"name": "required",
	}
}

func (r *Unauthorize) Messages() map[string]string {
	return map[string]string{}
}

func (r *Unauthorize) Attributes() map[string]string {
	return map[string]string{}
}

func (r *Unauthorize) PrepareForValidation(data validation.Data) {

}
