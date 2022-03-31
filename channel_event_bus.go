package waterlink

import "github.com/lukasl-dev/waterlink/event"

// ChannelEventBus is a EventBus implementation that uses Go channels to emit
// events.
type ChannelEventBus struct {
	playerUpdates   chan event.PlayerUpdate
	stats           chan event.Stats
	trackEnds       chan event.TrackEnd
	trackExceptions chan event.TrackException
	trackStucks     chan event.TrackStuck
	trackStarts     chan event.TrackStart
	webSocketCloses chan event.WebSocketClosed
}

var _ EventBus = (*ChannelEventBus)(nil)

// NewChannelEventBus creates a new ChannelEventBus and returns a pointer to it.
func NewChannelEventBus() *ChannelEventBus {
	return &ChannelEventBus{
		playerUpdates:   make(chan event.PlayerUpdate),
		stats:           make(chan event.Stats),
		trackEnds:       make(chan event.TrackEnd),
		trackExceptions: make(chan event.TrackException),
		trackStucks:     make(chan event.TrackStuck),
		trackStarts:     make(chan event.TrackStart),
		webSocketCloses: make(chan event.WebSocketClosed),
	}
}

func (c ChannelEventBus) PlayerUpdates() <-chan event.PlayerUpdate {
	return c.playerUpdates
}

func (c ChannelEventBus) EmitPlayerUpdate(e event.PlayerUpdate) {
	c.playerUpdates <- e
}

func (c ChannelEventBus) EmitStats(e event.Stats) {
	c.stats <- e
}

func (c ChannelEventBus) Stats() <-chan event.Stats {
	return c.stats
}

func (c ChannelEventBus) EmitTrackEnd(e event.TrackEnd) {
	c.trackEnds <- e
}

func (c ChannelEventBus) TrackEnds() <-chan event.TrackEnd {
	return c.trackEnds
}

func (c ChannelEventBus) EmitTrackException(e event.TrackException) {
	c.trackExceptions <- e
}

func (c ChannelEventBus) TrackExceptions() <-chan event.TrackException {
	return c.trackExceptions
}

func (c ChannelEventBus) EmitTrackStuck(e event.TrackStuck) {
	c.trackStucks <- e
}

func (c ChannelEventBus) TrackStucks() <-chan event.TrackStuck {
	return c.trackStucks
}

func (c ChannelEventBus) EmitTrackStart(e event.TrackStart) {
	c.trackStarts <- e
}

func (c ChannelEventBus) TrackStarts() <-chan event.TrackStart {
	return c.trackStarts
}

func (c ChannelEventBus) EmitWebSocketClosed(e event.WebSocketClosed) {
	c.webSocketCloses <- e
}

func (c ChannelEventBus) WebSocketCloses() <-chan event.WebSocketClosed {
	return c.webSocketCloses
}
