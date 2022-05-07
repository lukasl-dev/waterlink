package filter

type Karaoke struct {
	Level       float32 `json:"level,omitempty"`
	MonoLevel   float32 `json:"monoLevel,omitempty"`
	FilterBand  float32 `json:"filterBand,omitempty"`
	FilterWidth float32 `json:"filterWidth,omitempty"`
}
