package websocket

import (
	"encoding/json"
	"fmt"
	event2 "github.com/lukasl-dev/waterlink/v3/lavalink/api/websocket/event"
	message2 "github.com/lukasl-dev/waterlink/v3/lavalink/api/websocket/message"
)

// Router holds a set of channels to distribute incoming messages from
// the websocket connection.
type Router struct {
	ready        chan message2.Ready
	stats        chan message2.Stats
	playerUpdate chan message2.PlayerUpdate

	trackStart      chan event2.TrackStart
	trackEnd        chan event2.TrackEnd
	trackException  chan event2.TrackException
	trackStuck      chan event2.TrackStuck
	webSocketClosed chan event2.WebSocketClosed
}

// newMessageRouter returns a new Router.
func newMessageRouter() *Router {
	return &Router{
		ready:           make(chan message2.Ready),
		stats:           make(chan message2.Stats),
		playerUpdate:    make(chan message2.PlayerUpdate),
		trackStart:      make(chan event2.TrackStart),
		trackEnd:        make(chan event2.TrackEnd),
		trackException:  make(chan event2.TrackException),
		trackStuck:      make(chan event2.TrackStuck),
		webSocketClosed: make(chan event2.WebSocketClosed),
	}
}

// Ready returns a channel that receives a message.Ready as soon as one is
// received from the websocket connection.
func (ro *Router) Ready() <-chan message2.Ready {
	return ro.ready
}

// Stats returns a channel that receives a message.Stats as soon as one is
// received from the websocket connection.
func (ro *Router) Stats() <-chan message2.Stats {
	return ro.stats
}

// PlayerUpdate returns a channel that receives a message.PlayerUpdate as soon
// as one is received from the websocket connection.
func (ro *Router) PlayerUpdate() <-chan message2.PlayerUpdate {
	return ro.playerUpdate
}

// TrackStart returns a channel that receives a event.TrackStart as soon as one
// is received from the websocket connection.
func (ro *Router) TrackStart() <-chan event2.TrackStart {
	return ro.trackStart
}

// TrackEnd returns a channel that receives a event.TrackEnd as soon as one is
// received from the websocket connection.
func (ro *Router) TrackEnd() <-chan event2.TrackEnd {
	return ro.trackEnd
}

// TrackException returns a channel that receives a event.TrackException as soon
// as one is received from the websocket connection.
func (ro *Router) TrackException() <-chan event2.TrackException {
	return ro.trackException
}

// TrackStuck returns a channel that receives a event.TrackStuck as soon as one
// is received from the websocket connection.
func (ro *Router) TrackStuck() <-chan event2.TrackStuck {
	return ro.trackStuck
}

// WebSocketClosed returns a channel that receives a event.WebSocketClosed as
// soon as one is received from the websocket connection.
func (ro *Router) WebSocketClosed() <-chan event2.WebSocketClosed {
	return ro.webSocketClosed
}

// route processes incoming messages from the websocket connection and
// distributes them to the corresponding channel.
func (ro *Router) route(raw []byte) error {
	var base message2.Base
	if err := json.Unmarshal(raw, &base); err != nil {
		return fmt.Errorf("waterlink: websocket: router: %w", err)
	}

	switch base.Op {
	case message2.OpReady:
		var ready message2.Ready
		if err := json.Unmarshal(raw, &ready); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.ready <- ready

	case message2.OpStats:
		var stats message2.Stats
		if err := json.Unmarshal(raw, &stats); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.stats <- stats

	case message2.OpPlayerUpdate:
		var playerUpdate message2.PlayerUpdate
		if err := json.Unmarshal(raw, &playerUpdate); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.playerUpdate <- playerUpdate

	case message2.OpEvent:
		var evt message2.Event
		if err := json.Unmarshal(raw, &evt); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		return ro.routeEvent(raw, evt)
	}

	return nil
}

// routeEvent routes the given event message and distributes it in the respective
// event channel.
func (ro *Router) routeEvent(raw []byte, evt message2.Event) error {
	switch evt.Type {
	case event2.TypeTrackStart:
		var trackStart event2.TrackStart
		if err := json.Unmarshal(raw, &trackStart); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.trackStart <- trackStart

	case event2.TypeTrackEnd:
		var trackEnd event2.TrackEnd
		if err := json.Unmarshal(raw, &trackEnd); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.trackEnd <- trackEnd

	case event2.TypeTrackException:
		var trackException event2.TrackException
		if err := json.Unmarshal(raw, &trackException); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.trackException <- trackException

	case event2.TypeTrackStuck:
		var trackStuck event2.TrackStuck
		if err := json.Unmarshal(raw, &trackStuck); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.trackStuck <- trackStuck

	case event2.TypeWebSocketClosed:
		var webSocketClosed event2.WebSocketClosed
		if err := json.Unmarshal(raw, &webSocketClosed); err != nil {
			return fmt.Errorf("waterlink: websocket: router: %w", err)
		}

		ro.webSocketClosed <- webSocketClosed
	}

	return nil
}
