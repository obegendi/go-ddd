package pkg

type EventHandler interface {
	Notify(event Event)
}

type EventPublisher struct {
	handlers map[string][]EventHandler
}

func (e *EventPublisher) Subscribe(handler EventHandler, events ...Event) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

func (e *EventPublisher) Notify(event Event) {
	for _, handler := range e.handlers[event.Name()] {
		handler.Notify(event)
	}
}
