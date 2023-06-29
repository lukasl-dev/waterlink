package filter

type Vibrato struct {
	// Frequency is the frequency. Range is 0.0 to 14.0.
	Frequency float32 `json:"frequency,omitempty"`

	// Depth is the vibrato depth. Range is 0.0 to 1.0
	Depth float32 `json:"depth,omitempty"`
}
