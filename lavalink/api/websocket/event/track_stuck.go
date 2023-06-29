package event

import (
	"encoding/json"
	"github.com/lukasl-dev/waterlink/v3/lavalink/player"
	"time"
)

// TrackStuck is emitted when a track gets stuck while playing.
type TrackStuck struct {
	// EncodedTrack is the base64-encoded track that got stuck.
	EncodedTrack player.EncodedTrack `json:"encodedTrack,omitempty"`

	// Threshold is the threshold that was exceeded.
	Threshold time.Duration `json:"threshold,omitempty"`
}

var (
	_ json.Marshaler   = (*TrackStuck)(nil)
	_ json.Unmarshaler = (*TrackStuck)(nil)
)

type trackStuck struct {
	EncodedTrack string `json:"encodedTrack,omitempty"`
	ThresholdMS  uint   `json:"thresholdMS,omitempty"`
}

func (t *TrackStuck) MarshalJSON() ([]byte, error) {
	return json.Marshal(trackStuck{
		EncodedTrack: string(t.EncodedTrack),
		ThresholdMS:  uint(t.Threshold.Milliseconds()),
	})
}

func (t *TrackStuck) UnmarshalJSON(bytes []byte) error {
	var u trackStuck
	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	t.EncodedTrack = player.EncodedTrack(u.EncodedTrack)
	t.Threshold = time.Duration(u.ThresholdMS) * time.Millisecond

	return nil
}
