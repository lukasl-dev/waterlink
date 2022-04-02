package waterlink

import (
	"strconv"
	"time"
)

// PlayParams contains optional parameters of the Guild.Play() method.
type PlayParams struct {
	// StartTime is the time in milliseconds to start playing the track at.
	StartTime time.Duration `json:"startTime,omitempty"`

	// EndTime is the time in milliseconds to end playing the track at.
	EndTime time.Duration `json:"endTime,omitempty"`

	// Volume is the new volume of the player. The value must be between 0 and
	// 1000. Defaults to 100.
	Volume uint8 `json:"volume,omitempty"`

	// NoReplace is the optional flag to not replace the current track.
	NoReplace bool `json:"noReplace,omitempty"`

	// Pause is the optional flag to pause the playback.
	Pause bool `json:"pause,omitempty"`
}

// defaultPlayParams are the default PlayParams for the Guild.Play() method.
var defaultPlayParams = PlayParams{}

func (p PlayParams) startTime() string {
	if p.StartTime == 0 {
		return ""
	}
	return strconv.Itoa(int(p.StartTime))
}

func (p PlayParams) endTime() string {
	if p.EndTime == 0 {
		return ""
	}
	return strconv.Itoa(int(p.EndTime))
}

func (p PlayParams) volume() string {
	if p.Volume == 0 {
		return ""
	}
	return strconv.Itoa(int(p.Volume))
}
