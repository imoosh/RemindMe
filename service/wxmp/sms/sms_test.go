package sms

import "testing"

func TestTencentSmsService_AddSmsTemplate(t *testing.T) {
    (&SmsService{}).AddSmsTemplate()
}

func TestSmsService_SendSmsMessage(t *testing.T) {
    (&SmsService{}).SendSmsMessage("18113007510", nil)
}
