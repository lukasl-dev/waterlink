package event

import (
	"github.com/lukasl-dev/waterlink/v3/lavalink/player"
)

// TrackEnd is emitted when a track ends.
type TrackEnd struct {
	// EncodedTrack is the base64-encoded track that ended playing.
	EncodedTrack player.EncodedTrack `json:"encodedTrack,omitempty"`

	// Reason is the reason the track ended.
	Reason TrackEndReason `json:"reason,omitempty"`
}

// TrackEndReason represents a possible reason why a track ended.
type TrackEndReason string

const (
	// TrackEndReasonFinished occurs when a track finished playing.
	TrackEndReasonFinished TrackEndReason = "FINISHED"

	// TrackEndReasonLoadFailed occurs when a track failed to load.
	TrackEndReasonLoadFailed TrackEndReason = "LOAD_FAILED"

	// TrackEndReasonStopped occurs when a track was stopped.
	TrackEndReasonStopped TrackEndReason = "STOPPED"

	// TrackEndReasonReplaced occurs when a track was replaced.
	TrackEndReasonReplaced TrackEndReason = "REPLACED"

	// TrackEndReasonCleanup occurs when a track was cleaned up.
	TrackEndReasonCleanup TrackEndReason = "CLEANUP"
)
