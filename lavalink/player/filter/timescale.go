package filter

type Timescale struct {
	// Speed is the playback speed. Must be higher than or equal to 0.0. Defaults
	// to 1.0.
	Speed float32 `json:"speed,omitempty"`

	// Pitch is the pitch. Must be higher than or equal to 0.0 Defaults to 1.0.
	Pitch float32 `json:"pitch,omitempty"`

	// Rate is the rate. Must be higher than or equal to 0.0. Defaults to 1.0.
	Rate float32 `json:"rate,omitempty"`
}
