package helper

import "github.com/gookit/validate"

func Validate(value interface{}, prefix string) error {
	v := validate.Struct(value)
	v = v.WithMessages(map[string]string{
		"uint":      prefix + ".non_integer_{field}",
		"state":     prefix + ".invalid_{field}",
		"role":      prefix + ".invalid_{field}",
		"email":     prefix + ".invalid_{field}",
		"password":  prefix + ".invalid_{field}",
		"required":  prefix + ".invalid_{field}",
		"userState": prefix + ".invalid_{field}",
		"userRole":  prefix + ".invalid_{field}",
	})
}
