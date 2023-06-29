package player

import (
	"encoding/json"
	"time"
)

type Update struct {
	// values is a map of values that should be updated.
	values map[string]any
}

var _ json.Marshaler = (*Update)(nil)

func NewUpdate() *Update {
	return &Update{values: make(map[string]any)}
}

func (u *Update) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.values)
}

// EncodedTrack defines the encoded track with which the player should be updated
// with. If empty, the player stops playing.
func (u *Update) EncodedTrack(encodedTrack string) *Update {
	_, hasIdentifier := u.values["identifier"]
	if hasIdentifier {
		panic("waterlink: player: cannot set encoded track and identifier at the same time")
	}

	if encodedTrack == "" {
		u.values["encodedTrack"] = nil
	} else {
		u.values["encodedTrack"] = encodedTrack
	}

	return u
}

// Identifier defines the identifier of the track to play.
func (u *Update) Identifier(identifier string) *Update {
	_, hasEncodedTrack := u.values["encodedTrack"]
	if hasEncodedTrack {
		panic("waterlink: player: cannot set encoded track and identifier at the same time")
	}

	u.values["identifier"] = identifier
	return u
}

// Position defines the track position.
func (u *Update) Position(position time.Duration) *Update {
	u.values["position"] = position.Milliseconds()
	return u
}

// EndTime defines the track end time.
func (u *Update) EndTime(endTime time.Duration) *Update {
	u.values["endTime"] = endTime.Milliseconds()
	return u
}

// Volume defines the volume of the player.
func (u *Update) Volume(volume uint) *Update {
	u.values["volume"] = volume
	return u
}

// Paused defines whether the player is paused.
func (u *Update) Paused(paused bool) *Update {
	u.values["paused"] = paused
	return u
}

// Filters defines the new filters to apply. This will override all previously
// applied filters.
func (u *Update) Filters(filters Filters) *Update {
	u.values["filters"] = filters
	return u
}

// Voice defines information required for connecting to Discord, without
// connected or ping.
func (u *Update) Voice(voice VoiceState) *Update {
	u.values["voice"] = voice
	return u
}
