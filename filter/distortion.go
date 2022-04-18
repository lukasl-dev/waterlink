package filter

type Distortion struct {
	SinOffset float32 `json:"sinOffset,omitempty"`
	SinScale  float32 `json:"sinScale,omitempty"`
	CosOffset float32 `json:"cosOffset,omitempty"`
	CosScale  float32 `json:"cosScale,omitempty"`
	TanOffset float32 `json:"tanOffset,omitempty"`
	TanScale  float32 `json:"tanScale,omitempty"`
	Offset    float32 `json:"offset,omitempty"`
	Scale     float32 `json:"scale,omitempty"`
}
