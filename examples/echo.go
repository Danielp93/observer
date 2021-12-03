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

	// Subscribe Listener to Notifier
	n.Subscribe(l)
	defer n.Unsubscribe(l)

	m := observer.NewMessage("ECHO", "Hello from the other side!")

	n.Notify(m)
}
