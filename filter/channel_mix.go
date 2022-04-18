package filter

type ChannelMix struct {
	LeftToLeft   float32 `json:"leftToLeft,omitempty"`
	LeftToRight  float32 `json:"leftToRight,omitempty"`
	RightToLeft  float32 `json:"rightToLeft,omitempty"`
	RightToRight float32 `json:"rightToRight,omitempty"`
}
