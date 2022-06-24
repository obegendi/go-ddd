package main

type EventBus interface {
	PublishEvent(EventMessage)
	AddHandler(EventHandler, ...interface{})
}

// InternalEventBus provides a lightweight in process event bus
type InternalEventBus struct {
	eventHandlers map[string]map[EventHandler]struct{}
}

// NewInternalEventBus constructs a new InternalEventBus
func NewInternalEventBus() *InternalEventBus {
	b := &InternalEventBus{
		eventHandlers: make(map[string]map[EventHandler]struct{}),
	}
	return b
}

type EventHandler interface {
	Handle(EventMessage)
}

type EventMessage interface {

	// AggregateID returns the ID of the Aggregate that the event relates to
	AggregateID() string

	// GetHeaders returns the key value collection of headers for the event.
	//
	// Headers are metadata about the event that do not form part of the
	// actual event but are still required to be persisted alongside the event.
	GetHeaders() map[string]interface{}

	// SetHeader sets the value of the header specified by the key
	SetHeader(string, interface{})

	// Returns the actual event which is the payload of the event message.
	Event() interface{}

	// EventType returns a string descriptor of the command name
	EventType() string

	// Version returns the version of the event
	Version() *int
}
