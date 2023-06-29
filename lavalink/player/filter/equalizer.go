package filter

type EqualizerBand struct {
	// Band is the band number. Range is 0-14.
	Band uint `json:"band,omitempty"`

	// Gain is the multiplier for the band. Range is -0.25 to 1.0, where -0.25
	// means the given band is completely muted, and 0.25 means it is doubled.
	// Defaults to 0.0.
	Gain float64 `json:"gain,omitempty"`
}
