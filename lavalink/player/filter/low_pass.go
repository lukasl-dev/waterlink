package filter

type LowPass struct {
	// Smoothing is the smoothing factor. Must be higher than or equal to 1.0.
	Smoothing float32 `json:"smoothing,omitempty"`
}
