package loadtrack

type Response struct {
	LoadType     LoadType     `json:"loadType,omitempty"`
	PlaylistInfo PlaylistInfo `json:"playlistInfo,omitempty"`
	Tracks       []Track      `json:"tracks,omitempty"`
	Exception    Exception    `json:"exception,omitempty"`
}
