package payload

import "github.com/go-playground/validator/v10"

type CallbackRequest struct {
	Code  string `json:"code" query:"code" validate:"required"`
	State string `json:"state" query:"state"`
}

func (c *CallbackRequest) Validate() error {
	var validate = validator.New()
	return validate.Struct(c)
}
