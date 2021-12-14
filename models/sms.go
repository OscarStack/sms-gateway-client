package models

import (
	"strings"
)

type Sms struct {
	MobileNo string
	Message  string
}

type SmsNumber string

func (x *SmsNumber) Validate() bool {
	if !strings.HasPrefix(string(*x), "00") {
		return false
	}
	return true
}
