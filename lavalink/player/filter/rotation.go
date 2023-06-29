package filter

type Rotation struct {
	// RotationHz is the frequency of the audio rotation around the listener in
	// Hz. 0.2.
	RotationHz float32 `json:"rotationHz,omitempty"`
}
