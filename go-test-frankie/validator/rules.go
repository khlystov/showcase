package validator

import (
	"go-test-frankie/session"
	"regexp"
	"strings"

	vendorV "gopkg.in/go-playground/validator.v9"
)

// UniqueKey is in charge to validate that
// the list of "Keys" provided are unique to the call (no double-ups)
func UniqueKey(fl vendorV.FieldLevel) bool {
	var vals = make(map[string]bool)

	field := fl.Field()
	for i := 0; i < field.Len(); i++ {
		val := field.Index(i).FieldByName("KvpKey").String()
		_, ok := vals[val]
		if ok {
			return false
		}
		vals[val] = true
	}
	return true
}

// KvpType custom validator for KvpType field
func KvpType(fl vendorV.FieldLevel) bool {
	var re = regexp.MustCompile(KvpTypeRegexString)
	return re.Match([]byte(fl.Field().String()))
}

// ValidRaw custom validator for kvpValue field if kvpType is raw.*
// All raw.* fields will be base64 encoded so as to not interfere
// with JSON structuring. These are useful for returning/storing
// large quantities of data that doesn't necessarily require
// processing now, or may be useful to a calling client.
func ValidRaw(fl vendorV.FieldLevel) bool {
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
func UniqueSession(fl vendorV.FieldLevel) bool {
	val := fl.Field().String()
	ok := session.Store.Check(val)

	if ok {
		return false
	}

	go session.Store.Set(val)

	return true
}
