package filter

type LowPass struct {
	Smoothing float32 `json:"smoothing,omitempty"`
}
