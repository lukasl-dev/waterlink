package filter

type ChannelMix struct {
	// LeftToLeft is the left-to-left channel mix factor. Range is 0.0 to 1.0.
	LeftToLeft float32 `json:"leftToLeft,omitempty"`

	// LeftToRight is the left-to-right channel mix factor. Range is 0.0 to 1.0.
	LeftToRight float32 `json:"leftToRight,omitempty"`

	// RightToLeft is the right-to-left channel mix factor. Range is 0.0 is 1.0.
	RightToLeft float32 `json:"rightToLeft,omitempty"`

	// RightToRight is the right-to-right channel mix factor. Range is 0.0 to
	// 1.0.
	RightToRight float32 `json:"rightToRight,omitempty"`
}
