package main

import (
	"fmt"

	"github.com/danielp93/observer"
)

func main() {
	// Create Notifier
	n := observer.NewDefaultNotifier()
	if n == nil {
		return
	}
	n2 := observer.NewDefaultNotifier()
	if n == nil {
		return
	}

	// Create Listener from Listenerfunc (func(Message))
	// This is a simple Listener that will print out the message
	// If message is a Simplemessage it also will print the type and timestamp
	l := observer.ListenerFunc(func(message observer.Message) {
		switch m := message.(type) {
		case *observer.SimpleMessage:
			fmt.Printf("Received Message [%s] %s @%v\n", m.Type, m.Message, m.Timestamp)
		default:
			fmt.Println(message)
		}
	})
	// listener uses a channel under the hood, Close it for garbage collection
	defer l.Close()

	// Subscribe Listener to Notifier
	n.Subscribe(l)
	n2.Subscribe(l)

	m := observer.NewMessage("Hello from the other side!", "ECHO")
	m2 := observer.NewMessage("Hello also from me!", "OTHERECHO")

	n.Notify(m)
	n2.Notify(m2)
}
