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
	for _, l := range listeners {
		d.listeners = append(d.listeners, l)
		go l.Listen()
	}
}

func (d *DefaultNotifier) Unsubscribe(listeners ...Listener) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, l := range listeners {
		for i, value := range d.listeners {
			if value == l {
				d.listeners = append(d.listeners[:i], d.listeners[i+1:]...)
				close(l.GetChan())
			}
		}
	}
}

func (d *DefaultNotifier) Notify(m Message) {
	for _, l := range d.listeners {
		l.GetChan() <- m
	}
}
