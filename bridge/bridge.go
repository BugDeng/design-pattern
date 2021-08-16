// Package bridge decouple abstract and implementer
package bridge

import "fmt"

// Messager abstract api
type Messager interface {
	SendMessage(msg, target string) error
}

// Sender implementer api
type Sender interface {
	Send(msg, target string) error
}

// SMS ...
type SMS struct{}

// NewSMS ...
func NewSMS() Sender {
	return &SMS{}
}

// Send via SMS
func (sms *SMS) Send(msg, target string) error {
	fmt.Printf("send %s to %s via sms\n", msg, target)
	return nil
}

// Email ...
type Email struct{}

// NewEmail ...
func NewEmail() Sender {
	return &Email{}
}

// Send via Email
func (email *Email) Send(msg, target string) error {
	fmt.Printf("send %s to %s via email\n", msg, target)
	return nil
}

// Info level
type Info struct {
	sender Sender
}

// NewInfo initial by given sender
func NewInfo(s Sender) Messager {
	return &Info{sender: s}
}

// SendMessage send given msg to target with info level
func (info *Info) SendMessage(msg, target string) error {
	return info.sender.Send(fmt.Sprintf("[Info] %s", msg), target)
}

// Urgency level
type Urgency struct {
	sender Sender
}

// NewUrgency initial by given sender
func NewUrgency(s Sender) Messager {
	return &Urgency{sender: s}
}

// SendMessage send given msg to target with info level
func (urgency *Urgency) SendMessage(msg, target string) error {
	return urgency.sender.Send(fmt.Sprintf("[Urgency] %s", msg), target)
}
