package filter

type EqualizerBand struct {
	Band uint8   `json:"band,omitempty"`
	Gain float32 `json:"gain,omitempty"`
}
