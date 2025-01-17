package bot

import "github.com/disgoorg/disgo/discord"

// DefaultEventManagerConfig returns a new EventManagerConfig with all default values.
func DefaultEventManagerConfig() *EventManagerConfig {
	return &EventManagerConfig{}
}

// EventManagerConfig can be used to configure the EventManager.
type EventManagerConfig struct {
	EventListeners     []EventListener
	RawEventsEnabled   bool
	AsyncEventsEnabled bool

	GatewayHandlers   map[discord.GatewayEventType]GatewayEventHandler
	HTTPServerHandler HTTPServerEventHandler
}

// EventManagerConfigOpt is a functional option for configuring an EventManager.
type EventManagerConfigOpt func(config *EventManagerConfig)

// Apply applies the given EventManagerConfigOpt(s) to the EventManagerConfig.
func (c *EventManagerConfig) Apply(opts []EventManagerConfigOpt) {
	for _, opt := range opts {
		opt(c)
	}
}

// WithListeners adds the given EventListener(s) to the EventManagerConfig.
func WithListeners(listeners ...EventListener) EventManagerConfigOpt {
	return func(config *EventManagerConfig) {
		config.EventListeners = append(config.EventListeners, listeners...)
	}
}

// WithListenerFunc adds the given ListenerFunc(s) to the EventManagerConfig.
func WithListenerFunc[E Event](listenerFunc func(e E)) EventManagerConfigOpt {
	return WithListeners(NewListenerFunc(listenerFunc))
}

// WithRawEventsEnabled enables/disables the raw events.
func WithRawEventsEnabled() EventManagerConfigOpt {
	return func(config *EventManagerConfig) {
		config.RawEventsEnabled = true
	}
}

// WithAsyncEventsEnabled enables/disables the async events.
func WithAsyncEventsEnabled() EventManagerConfigOpt {
	return func(config *EventManagerConfig) {
		config.AsyncEventsEnabled = true
	}
}

// WithGatewayHandlers overrides the default GatewayEventHandler(s) in the EventManagerConfig.
func WithGatewayHandlers(handlers map[discord.GatewayEventType]GatewayEventHandler) EventManagerConfigOpt {
	return func(config *EventManagerConfig) {
		config.GatewayHandlers = handlers
	}
}

// WithHTTPServerHandler overrides the given HTTPServerEventHandler in the EventManagerConfig.
func WithHTTPServerHandler(handler HTTPServerEventHandler) EventManagerConfigOpt {
	return func(config *EventManagerConfig) {
		config.HTTPServerHandler = handler
	}
}
