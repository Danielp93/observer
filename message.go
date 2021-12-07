package observer

import (
	"fmt"
	"time"
)

type Message interface {
	fmt.Stringer
	Type() string
	Timestamp() time.Time
}

type SimpleMessage struct {
	ts time.Time
	t  string
	m  string
}

func NewMessage(message string, messagetype string) Message {
	return &SimpleMessage{
		ts: time.Now(),
		t:  messagetype,
		m:  message,
	}
}

func (s *SimpleMessage) String() string {
	return s.m
}

func (s *SimpleMessage) Timestamp() time.Time {
	return s.ts
}
func (s *SimpleMessage) Type() string {
	return s.t
}
