package message

// Event represents an incoming event message which only contains base
// information about the event.
type Event struct {
	Guild

	// Type determines the type of this event.
	Type EventType `json:"type"`
}

// EventType determines the type of incoming events.
type EventType string

const (
	EventTypeTrackStart      EventType = "TrackStartEvent"
	EventTypeTrackEnd        EventType = "TrackEndEvent"
	EventTypeTrackException  EventType = "TrackExceptionEvent"
	EventTypeTrackStuck      EventType = "TrackStuckEvent"
	EventTypeWebSocketClosed EventType = "WebSocketClosedEvent"
)
