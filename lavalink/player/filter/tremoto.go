package filter

type Tremolo struct {
	// Frequency is the frequency. Must be higher than or equal to 0.0.
	Frequency float32 `json:"frequency,omitempty"`

	// Depth is the tremolo depth. Must be higher than or equal 0.0.
	Depth float32 `json:"depth,omitempty"`
}
