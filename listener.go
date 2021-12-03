package observer

import (
	"sync"
)

type Listener interface {
	Listen()
	Send(Message)

	Close()
}

type ListenFunc func(Message)

func ListenerFunc(l ListenFunc) Listener {
	listener := &DefaultListener{
		mChan: make(chan Message),
		f:     l,
		wg:    &sync.WaitGroup{},
	}
	listener.Listen()

	return listener
}

type DefaultListener struct {
	mChan chan Message
	f     ListenFunc
	wg    *sync.WaitGroup
}

func (l *DefaultListener) Send(message Message) {
	l.mChan <- message
}

func (l *DefaultListener) Listen() {
	l.wg.Add(1)
	go func() {
		for message := range l.mChan {
			l.f(message)
		}
		l.wg.Done()
	}()
}

func (l *DefaultListener) Close() {
	close(l.mChan)
	l.wg.Wait()
}
