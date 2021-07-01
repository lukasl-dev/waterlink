package loadtrack

type Track struct {
	ID   string    `json:"track,omitempty"`
	Info TrackInfo `json:"info,omitempty"`
}
