package message

import "github.com/lukasl-dev/waterlink/v2/filter"

type Filters struct {
	Incoming
	Guild
	Volume     float32           `json:"volume,omitempty"`
	Karaoke    filter.Karaoke    `json:"karaoke,omitempty"`
	Timescale  filter.Timescale  `json:"timescale,omitempty"`
	Tremolo    filter.Tremolo    `json:"tremolo,omitempty"`
	Vibrato    filter.Vibrato    `json:"vibrato,omitempty"`
	Rotation   filter.Rotation   `json:"rotation,omitempty"`
	Distortion filter.Distortion `json:"distortion,omitempty"`
	ChannelMix filter.ChannelMix `json:"channelMix,omitempty"`
	LowPass    filter.LowPass    `json:"lowPass,omitempty"`
}
