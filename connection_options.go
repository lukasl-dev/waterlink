package waterlink

// ConnectionOptions contains optional values of a Connection.
type ConnectionOptions struct {
	// EventHandler is the EventHandler to handle incoming events. If nil, the
	// connection will not listen for incoming events.
	EventHandler EventHandler `json:"eventHandler,omitempty"`

	// HandleEventError is the function to be called when an error occurs while
	// listening for incoming events.
	HandleEventError func(err error)
}

// defaultConnectionOptions are the default ConnectionOptions to use if none are
// provided.
var defaultConnectionOptions = ConnectionOptions{}
