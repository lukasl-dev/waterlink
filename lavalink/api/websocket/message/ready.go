package message

// Ready is dispatched by Lavalink upon successful connection and authorization.
type Ready struct {
	// Resumed is whether a session was resumed.
	Resumed bool `json:"resumed"`

	// SessionID is the Lavalink session ID of this connection. Not to be
	// confused with a Discord voice session ID.
	SessionID string `json:"sessionId,omitempty"`
}
