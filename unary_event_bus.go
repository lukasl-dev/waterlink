package waterlink

import (
	"github.com/lukasl-dev/waterlink/event"
)

// unaryEventBus is used when there is only one subscriber that generalises all
// incoming events.
type unaryEventBus struct {
	// fn is the function that will be called when any event is received.
	fn func(e interface{})
}

var _ EventBus = (*unaryEventBus)(nil)

// NewUnaryEventBus returns a EventBus that will call fn() when any event is
// emitted.
func NewUnaryEventBus(fn func(e interface{})) EventBus {
	return &unaryEventBus{fn: fn}
}

func (u unaryEventBus) EmitPlayerUpdate(e event.PlayerUpdate) {
	u.fn(e)
}

func (u unaryEventBus) EmitStats(e event.Stats) {
	u.fn(e)
}

func (u unaryEventBus) EmitTrackEnd(e event.TrackEnd) {
	u.fn(e)
}

func (u unaryEventBus) EmitTrackException(e event.TrackException) {
	u.fn(e)
}

func (u unaryEventBus) EmitTrackStart(e event.TrackStart) {
	u.fn(e)
}

func (u unaryEventBus) EmitTrackStuck(e event.TrackStuck) {
	u.fn(e)
}

func (u unaryEventBus) EmitWebSocketClosed(e event.WebSocketClosed) {
	u.fn(e)
}
