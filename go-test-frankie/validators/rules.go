package validators

import (
	"go-test-frankie/sessions"
	"regexp"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// KvpType custom validator for KvpType field
func KvpType(fl validator.FieldLevel) bool {
	var re = regexp.MustCompile(KvpTypeRegexString)
	return re.Match([]byte(fl.Field().String()))
}

// ValidRaw custom validator for kvpValue field if kvpType is raw.*
// All raw.* fields will be base64 encoded so as to not interfere
// with JSON structuring. These are useful for returning/storing
// large quantities of data that doesn't necessarily require
// processing now, or may be useful to a calling client.
func ValidRaw(fl validator.FieldLevel) bool {
	kvpF := fl.Parent().FieldByName("KvpType")

	if kvpF.IsValid() && strings.Contains(kvpF.String(), "raw") {
		var re = regexp.MustCompile(Base64RegexString)
		return re.Match([]byte(fl.Field().String()))
	}

	return true
}

// UniqueSession custom validator for checkSessionKey
// The unique session based ID that will be checked against the service.
// Service key must be unique or an error will be returned.
// Session keys should be unique across all calls for the life of the
// running process. If you stop/start the process,
// the uniqueness check will reset.
func UniqueSession(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	ok := sessions.Store.Check(val)

	if ok {
		return false
	}

	go sessions.Store.Set(val)

	return true
}
