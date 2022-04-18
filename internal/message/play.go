package message

// Play is the message type of opcode.Play messages.
type Play struct {
	Outgoing
	Guild

	// Track is the ID of the track to play.
	Track string `json:"track"`

	// StartTime is the optional start time of the track in milliseconds.
	StartTime string `json:"startTime,omitempty"`

	// EndTime is the optional end time of the track in milliseconds.
	EndTime string `json:"endTime,omitempty"`

	// Volume is the optional volume of the track in percent. It must be between
	// 0 and 1000. Defaults to 100.
	Volume string `json:"volume,omitempty"`

	// NoReplace is the optional flag to not replace the current track.
	NoReplace bool `json:"noReplace,omitempty"`

	// Pause is the optional flag to pause the playback.
	Pause bool `json:"pause,omitempty"`
}
