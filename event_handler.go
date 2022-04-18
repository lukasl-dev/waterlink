package waterlink

// EventHandler is an interface that is responsible for handling incoming events
// provided by the server.
type EventHandler interface {
	// HandleEvent handles the incoming event. The following events are available:
	//    - event.PlayerUpdate
	//    - event.Stats
	//    - event.TrackEnd
	//    - event.TrackException
	//    - event.TrackStart
	//    - event.TrackStuck
	//    - event.WebSocketClosed
	HandleEvent(evt interface{})
}

// EventHandlerFunc is a function that implements the EventHandler interface.
type EventHandlerFunc func(evt interface{})

var _ EventHandler = EventHandlerFunc(nil)

// HandleEvent calls itself using the provided event.
func (fn EventHandlerFunc) HandleEvent(evt interface{}) {
	fn(evt)
}
