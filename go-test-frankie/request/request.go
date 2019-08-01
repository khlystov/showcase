package request

import (
	"go-test-frankie/validator"
)

// DeviceCheck is base structure for binding http request data
type DeviceCheck struct {
	Item []struct {
		CheckType       string `json:"checkType" validate:"required,oneof=DEVICE BIOMETRIC COMBO"`
		ActivityType    string `json:"activityType" validate:"required,oneof=SIGNUP LOGIN PAYMENT CONFIRMATION|startswith=_"`
		CheckSessionKey string `json:"checkSessionKey" validate:"required,usession"`
		ActivityData    []struct {
			KvpKey   string `json:"kvpKey" validate:"required"`
			KvpValue string `json:"kvpValue" validate:"required,raw"`
			KvpType  string `json:"kvpType" validate:"required,kvptype"`
		} `json:"activityData" validate:"required,ukey,dive,required"`
	} `validate:"required,dive,required"`
}

// Validate is in charge of be sure that DeviceCheck
// structure contains only correct data based on business rules
func (r *DeviceCheck) Validate() (bool, error) {
	for _, item := range r.Item {
		err := validator.Validate.Struct(item)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
