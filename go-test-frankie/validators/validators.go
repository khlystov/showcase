package validators

import (
	"gopkg.in/go-playground/validator.v9"
)

// Validate is exported validator instance
var Validate = initV()

func initV() *validator.Validate {
	Validate := validator.New()

	// Attach custom validation rules
	Validate.RegisterValidation("usession", UniqueSession)
	Validate.RegisterValidation("kvptype", KvpType)
	Validate.RegisterValidation("raw", ValidRaw)

	return Validate
}
