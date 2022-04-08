package waterlink

import (
	"fmt"
	"github.com/lukasl-dev/waterlink/event"
)

// FunctionalEventBus is a EventBus implementation that calls pre-registered
// functions for each event emitted.
type FunctionalEventBus struct {
	playerUpdates   []func(event.PlayerUpdate)
	stats           []func(event.Stats)
	trackEnds       []func(event.TrackEnd)
	trackExceptions []func(event.TrackException)
	trackStarts     []func(event.TrackStart)
	trackStucks     []func(event.TrackStuck)
	websocketCloses []func(event.WebSocketClosed)
}

var _ EventBus = (*FunctionalEventBus)(nil)

// NewFunctionalEventBus creates a new FunctionEventBus and returns a pointer to
// it.
func NewFunctionalEventBus() *FunctionalEventBus {
	return new(FunctionalEventBus)
}

// EmitPlayerUpdate calls every registered functions that listen event.PlayerUpdate
// sequentially with the given event.
func (f *FunctionalEventBus) EmitPlayerUpdate(e event.PlayerUpdate) {
	for _, fn := range f.playerUpdates {
		fn(e)
	}
}

// HandlePlayerUpdate registers a function to be called when an event.PlayerUpdate
// is emitted.
func (f *FunctionalEventBus) HandlePlayerUpdate(fn func(update event.PlayerUpdate)) {
	f.playerUpdates = append(f.playerUpdates, fn)
}

// EmitStats calls every registered functions that listen event.Stats
// sequentially with the given event.
func (f *FunctionalEventBus) EmitStats(e event.Stats) {
	for _, fn := range f.stats {
		fn(e)
	}
}

// HandleStats registers a function to be called when an event.Stats is emitted.
func (f *FunctionalEventBus) HandleStats(fn func(e event.Stats)) {
	f.stats = append(f.stats, fn)
}

// EmitTrackEnd calls every registered functions that listen event.TrackEnd
// sequentially with the given event.
func (f *FunctionalEventBus) EmitTrackEnd(e event.TrackEnd) {
	for _, fn := range f.trackEnds {
		fn(e)
	}
}

// HandleTrackEnd calls every registered functions that listen event.TrackEnd
// sequentially with the given event.
func (f *FunctionalEventBus) HandleTrackEnd(fn func(e event.TrackEnd)) {
	f.trackEnds = append(f.trackEnds, fn)
}

// EmitTrackException calls every registered functions that listen
// event.TrackException sequentially with the given event.
func (f *FunctionalEventBus) EmitTrackException(e event.TrackException) {
	for _, fn := range f.trackExceptions {
		fn(e)
	}
}

// HandleTrackException registers a function to be called when an
// event.TrackException is emitted.
func (f *FunctionalEventBus) HandleTrackException(fn func(e event.TrackException)) {
	f.trackExceptions = append(f.trackExceptions, fn)
}

// EmitTrackStart calls every registered functions that listen event.TrackStart
// sequentially with the given event.
func (f *FunctionalEventBus) EmitTrackStart(e event.TrackStart) {
	fmt.Println("Emitting track start:", e)
	for _, fn := range f.trackStarts {
		fn(e)
	}
}

// HandleTrackStart registers a function to be called when an event.TrackStart
// is emitted.
func (f *FunctionalEventBus) HandleTrackStart(fn func(e event.TrackStart)) {
	f.trackStarts = append(f.trackStarts, fn)
}

// EmitTrackStuck calls every registered functions that listen event.TrackStuck
// sequentially with the given event.
func (f *FunctionalEventBus) EmitTrackStuck(e event.TrackStuck) {
	for _, fn := range f.trackStucks {
		fn(e)
	}
}

// HandleTrackStuck registers a function to be called when an event.TrackStuck
// is emitted.
func (f *FunctionalEventBus) HandleTrackStuck(fn func(e event.TrackStuck)) {
	f.trackStucks = append(f.trackStucks, fn)
}

// EmitWebSocketClosed calls every registered functions that listen
// event.WebSocketClosed sequentially with the given event.
func (f *FunctionalEventBus) EmitWebSocketClosed(e event.WebSocketClosed) {
	for _, fn := range f.websocketCloses {
		fn(e)
	}
}

// HandleWebSocketClosed registers a function to be called when an
// event.WebSocketClosed is emitted.
func (f *FunctionalEventBus) HandleWebSocketClosed(fn func(e event.WebSocketClosed)) {
	f.websocketCloses = append(f.websocketCloses, fn)
}
