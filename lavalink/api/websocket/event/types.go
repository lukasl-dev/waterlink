package event

// Type represents the type of events. It is used to distinguish between
// different events.
type Type string

const (
	// TypeTrackStart is emitted when a track starts playing.
	TypeTrackStart Type = "TrackStart"

	// TypeTrackEnd is emitted when a track ends.
	TypeTrackEnd Type = "TrackEndEvent"

	// TypeTrackException is emitted when a track throws an exception.
	TypeTrackException Type = "TrackExceptionEvent"

	// TypeTrackStuck is emitted when a track gets stuck while playing.
	TypeTrackStuck Type = "TrackStuckEvent"

	// TypeWebSocketClosed is emitted when the websocket connection to Discord
	// voice servers is closed.
	TypeWebSocketClosed Type = "WebSocketClosedEvent"
)
