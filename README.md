# Simple Observer (go)

A small simple library to use for sending around messages in-memory. It uses a notifier/listener style messaging.


## Installation

```shell
go get -u github.com/danielp93/observer
```
## Usage

* Create notifiers by implementing the `Notifier` interface (or spawn a provided `DefaultNotifier`)
* Create listeners by implementing the `Listener` interface (or spawn a provided `DefaultListener` by calling ListenerFunc(ListenFunc))
* A ListenerFunc is a function that takes a `Message` and calls that function whenever the `Listener` is `Notified`
* Notify `Listeners` by subscribing them to `Notifiers` and calling the `Notify(Message)` method receiver on the Notifier

* For advanced customization the library provides interfaces

* Be aware that this 


##
```Go
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

	// Subscribe Listener to Notifiers
	n.Subscribe(l)
	n2.Subscribe(l)

	m := observer.NewMessage("Hello from the other side!", "ECHO")
	m2 := observer.NewMessage("Hello also from me!", "ECHO2")

	n.Notify(m)
	n2.Notify(m2)
}
```

```shell
$ go run observer/examples/echo.go

Received Message [ECHO] Hello from the other side! @2021-12-03 21:08:01.5650042 +0100 CET m=+0.000041501
Received Message [OTHERECHO] Hello also from me! @2021-12-03 21:08:01.5650054 +0100 CET m=+0.000042601
```
