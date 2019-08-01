package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	body     string
	status   int
	response string
}{
	{
		`[ { "checkType": "DEVICE", "activityType": "SIGNUP", "checkSessionKey": "1", "activityData": [ { "kvpKey": "die", "kvpValue": "c2RmdZndlIHdlciB3ZWZl", "kvpType": "general.string" } ] } ]`,
		200,
		`{"puppy":true}`,
	},
	{
		`[ { "checkType": "DEVICE", "activityType": "SIGNUP", "checkSessionKey": "1", "activityData": [ { "kvpKey": "die", "kvpValue": "c2RmdZndlIHdlciB3ZWZl", "kvpType": "general.string" } ] } ]`,
		500,
		`{"code":0,"message":"Key: 'CheckSessionKey' Error:Field validation for 'CheckSessionKey' failed on the 'usession' tag"}`,
	},
	{
		`[ { "checkType": "DEVICE_LESS", "activityType": "SIGNUP", "checkSessionKey": "1", "activityData": [ { "kvpKey": "die", "kvpValue": "c2RmdZndlIHdlciB3ZWZl", "kvpType": "general.string" } ] } ]`,
		500,
		``,
	},
	{
		`[ { "checkType": "DEVICE_LESS", "activityType": "SIGNUP_LESS", "checkSessionKey": "1", "activityData": [ { "kvpKey": "die", "kvpValue": "c2RmdZndlIHdlciB3ZWZl", "kvpType": "general.string" } ] } ]`,
		500,
		``,
	},
	{
		`[ { "checkType": "DEVICE_LESS", "activityType": "SIGNUP_LESS", "checkSessionKey": "1", "activityData": [ { "kvpKey": "die", "kvpValue": "c2RmdZndlIHdlciB3ZWZl", "kvpType": "badtype" } ] } ]`,
		500,
		``,
	},
	{
		``,
		500,
		`{"code":0,"message":"Everything is wrong. Go fix it."}`,
	},
}

func TestIsGood(t *testing.T) {
	r := gin.Default()
	r.POST("/isgood", IsGood)

	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer([]byte(test.body)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, test.status, w.Code)
		if test.response != "" {
			assert.Equal(t, test.response, w.Body.String())
		}
	}
}
