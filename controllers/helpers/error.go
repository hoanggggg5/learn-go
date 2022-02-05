package helper

import (
	"errors"

	"github.com/gookit/validate"
)

func Validate(value interface{}, prefix string) error {
	v := validate.Struct(value)
	v = v.WithMessages(map[string]string{
		"email":    prefix + ".invalid_{field}",
		"password": prefix + ".invalid_{field}",
	})

	if !v.Validate() {
		for _, errs := range v.Errors.All() {
			for _, err := range errs {
				return errors.New(err)
			}
		}
	}
	return nil
}
