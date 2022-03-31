package waterlink

// ConnectionOptions contains optional values of a Connection.
type ConnectionOptions struct {
	// EventBus is the EventBus to emit incoming events to. If nil, the connection
	// will not listen for incoming events.
	EventBus EventBus `json:"eventBus,omitempty"`
}

// defaultConnectionOptions are the default ConnectionOptions to use if none are
// provided.
var defaultConnectionOptions = ConnectionOptions{}
