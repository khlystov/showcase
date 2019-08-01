package validator

import (
	vendorV "gopkg.in/go-playground/validator.v9"
)

// Validate is exported validator instance
var Validate = initV()

func initV() *vendorV.Validate {
	Validate := vendorV.New()

	// Attach custom validation rules
	Validate.RegisterValidation("usession", UniqueSession)
	Validate.RegisterValidation("kvptype", KvpType)
	Validate.RegisterValidation("raw", ValidRaw)
	Validate.RegisterValidation("ukey", UniqueKey)

	return Validate
}
