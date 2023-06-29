package filter

type Distortion struct {
	// SinOffset is the sin offset.
	SinOffset float32 `json:"sinOffset,omitempty"`

	// SinScale is the sin scale.
	SinScale float32 `json:"sinScale,omitempty"`

	// CosOffset is the cos offset.
	CosOffset float32 `json:"cosOffset,omitempty"`

	// CosScale is the cos scale.
	CosScale float32 `json:"cosScale,omitempty"`

	// TanOffset is the tan offset.
	TanOffset float32 `json:"tanOffset,omitempty"`

	// TanScale is the tan scale.
	TanScale float32 `json:"tanScale,omitempty"`

	// Offset is the offset.
	Offset float32 `json:"offset,omitempty"`

	// Scale is the scale.
	Scale float32 `json:"scale,omitempty"`
}
