package requests

import (
	"go-test-frankie/validators"
)

// DeviceCheck is base structure for binding http request data
type DeviceCheck []struct {
	CheckType       string `json:"checkType" validate:"required,oneof=DEVICE BIOMETRIC COMBO"`
	ActivityType    string `json:"activityType" validate:"required,oneof=SIGNUP LOGIN PAYMENT CONFIRMATION|startswith=_"`
	CheckSessionKey string `json:"checkSessionKey" validate:"required,usession"`
	ActivityData    []struct {
		KvpKey   string `json:"kvpKey" validate:"required"`
		KvpValue string `json:"kvpValue" validate:"required,raw"`
		KvpType  string `json:"kvpType" validate:"required,kvptype"`
	} `json:"activityData" validate:"required,dive,required"`
}

// Validate is in charge of be sure that DeviceCheck
// structure contains only correct data based on business rules
func (r *DeviceCheck) Validate() (bool, error) {
	for _, item := range *r {
		err := validators.Validate.Struct(item)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
