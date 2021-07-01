package loadtrack

type TrackInfo struct {
	Identifier string `json:"identifier,omitempty"`
	Seekable   bool   `json:"isSeekable,omitempty"`
	Author     string `json:"author,omitempty"`
	Length     uint   `json:"length,omitempty"`
	Stream     bool   `json:"isStream,omitempty"`
	Position   uint   `json:"position,omitempty"`
	Title      string `json:"title,omitempty"`
	URI        string `json:"uri,omitempty"`
}
