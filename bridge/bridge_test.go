package bridge

import "testing"

func TestBridge(t *testing.T) {
	sms := NewSMS()
	email := NewEmail()
	info := NewInfo(sms)
	info.SendMessage("haha", "10086")
	urgency := NewUrgency(email)
	urgency.SendMessage("crash", "a@a.mail")
}
