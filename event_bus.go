package waterlink

import "github.com/lukasl-dev/waterlink/event"

// EventBus is an interface that defines methods used to subscribe to incoming
// server messages.
type EventBus interface {
	// EmitPlayerUpdate emits incoming event.PlayerUpdate messages.
	EmitPlayerUpdate(e event.PlayerUpdate)

	// EmitStats emits incoming event.Stats messages.
	EmitStats(e event.Stats)

	// EmitTrackEnd emits incoming event.TrackEnd messages.
	EmitTrackEnd(e event.TrackEnd)

	// EmitTrackException emits incoming event.TrackException messages.
	EmitTrackException(e event.TrackException)

	// EmitTrackStart emits incoming event.TrackStart messages.
	EmitTrackStart(e event.TrackStart)

	// EmitTrackStuck emits incoming event.TrackStuck messages.
	EmitTrackStuck(e event.TrackStuck)

	// EmitWebSocketClosed emits incoming event.WebSocketClosed messages.
	EmitWebSocketClosed(e event.WebSocketClosed)
}
