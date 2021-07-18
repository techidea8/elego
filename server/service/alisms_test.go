package service

import "testing"

func TestSendNotifySms(t *testing.T) {
	err := SendNotifySms("15367151352", "来自湖南")
	if err != nil {
		t.Error(err)
	}
}

func TestSendRegisterSms(t *testing.T) {
	err := SendRegisterSms("15367151352", "湖南")
	if err != nil {
		t.Error(err)
	}
}
