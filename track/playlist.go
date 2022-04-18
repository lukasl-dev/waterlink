package track

// Playlist contains information about a playlist.
type Playlist struct {
	// Name is the name of the playlist.
	Name string `json:"name,omitempty"`

	// SelectedTrack is the index of the currently selected track. If no track
	// is selected, this is -1.
	SelectedTrack int `json:"selectedTrack,omitempty"`
}
