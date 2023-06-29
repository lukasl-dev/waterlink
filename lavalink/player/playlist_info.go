package player

type PlaylistInfo struct {
	// Name is the name of the loaded playlist.
	Name string `json:"name,omitempty"`

	// SelectedTrack is the index of the selected track in this playlist. If
	// none is selected, this will be -1.
	SelectedTrack int `json:"selectedTrack,omitempty"`
}

// IsPresent returns whether a playlist has been loaded.
func (pi PlaylistInfo) IsPresent() bool {
	return pi.Name != ""
}
