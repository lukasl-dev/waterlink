package player

type VoiceState struct {
	// Token is the Discord voice token to authenticate with.
	Token string `json:"token,omitempty"`

	// Endpoint is the Discord voice endpoint to connect to.
	Endpoint string `json:"endpoint,omitempty"`

	// SessionID is the Discord voice session ID to authenticate with.
	SessionID string `json:"sessionId,omitempty"`

	// Connected is whether the player is connected.
	Connected bool `json:"connected,omitempty"`

	// Ping is the round-trip latency in milliseconds to the voice gateway. If
	// not connected, this is -1.
	Ping int `json:"ping,omitempty"`
}
