package observer

type Listener interface {
	GetChan() chan Message
	Listen()
}

type ListenFunc func(Message)

type DefaultListener struct {
	l chan Message
	h ListenFunc
}

func (l *DefaultListener) GetChan() chan Message {
	return l.l
}

func (l *DefaultListener) Listen() {
	for {
		message := <-l.GetChan()
		l.h(message)
	}
}

func ListenerFunc(handlerFunc ListenFunc) Listener {
	return &DefaultListener{
		l: make(chan Message),
		h: handlerFunc,
	}
}
