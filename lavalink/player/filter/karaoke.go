package filter

type Karaoke struct {
	// Level is the level of the karaoke effect. Range is 0.0 to 1.0, where 0.0
	// is no effect and 1.0 is full effect.
	Level float32 `json:"level,omitempty"`

	// MonoLevel is the level of the mono effect. Range is 0.0 to 1.0, where 0.0
	// is no effect and 1.0 is full effect.
	MonoLevel float32 `json:"monoLevel,omitempty"`

	// FilterBand is the filter band.
	FilterBand float32 `json:"filterBand,omitempty"`

	// FilterWidth is the filter width.
	FilterWidth float32 `json:"filterWidth,omitempty"`
}
