package observer

import (
	"fmt"
	"time"
)

type Message interface {
	fmt.Stringer
}

type SimpleMessage struct {
	Timestamp time.Time
	Type      string
	Message   string
}

func NewMessage(message string, messagetype string) Message {
	return SimpleMessage{
		Timestamp: time.Now(),
		Type:      messagetype,
		Message:   message,
	}
}

func (s SimpleMessage) Timestamp() time.Time {
	return s.Ts
}

func (s SimpleMessage) Type() string {
	return s.T
}

func (s SimpleMessage) String() string {
	return s.Message
}
