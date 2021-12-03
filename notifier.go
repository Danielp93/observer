package observer

import (
	"sync"
)

type Notifier interface {
	Subscribe(l ...Listener)
	Unsubscribe(l ...Listener)
	Notify(Message)
}

type DefaultNotifier struct {
	listeners []Listener
	mu        *sync.Mutex
}

func NewDefaultNotifier() Notifier {
	return &DefaultNotifier{
		mu: &sync.Mutex{},
	}
}

func (d *DefaultNotifier) Subscribe(listeners ...Listener) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.listeners = append(d.listeners, listeners...)
}

func (d *DefaultNotifier) Unsubscribe(listeners ...Listener) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, l := range d.listeners {
		for i, listener := range listeners {
			if listener == l {
				d.listeners = append(d.listeners[:i], d.listeners[i+1:]...)
			}
		}
	}
}

func (d *DefaultNotifier) Notify(m Message) {
	for _, l := range d.listeners {
		l.Send(m)
	}
}
