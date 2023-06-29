package event

import (
	"github.com/lukasl-dev/waterlink/v3/lavalink"
	"github.com/lukasl-dev/waterlink/v3/lavalink/player"
)

// TrackException is emitted when a track throws an exception.
type TrackException struct {
	// EncodedTrack is the base64-encoded track that threw the exception.
	EncodedTrack player.EncodedTrack `json:"encodedTrack,omitempty"`

	// Exception is the occurred exception.
	Exception lavalink.Exception `json:"exception,omitempty"`
}
