package observer

import (
	"fmt"
	"time"
)

type Message interface {
	Type() string
	Timestamp() time.Time
	fmt.Stringer
}

type SimpleMessage struct {
	Ts      time.Time
	T       string
	Message string
}

func NewMessage(message string, messagetype string) Message {
	return SimpleMessage{
		Ts:      time.Now(),
		T:       messagetype,
		Message: message,
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
