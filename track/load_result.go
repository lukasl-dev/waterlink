package track

// LoadResult represents the result of a load tracks request/operation.
type LoadResult struct {
	// LoadType the type of the result.
	LoadType LoadType `json:"loadType,omitempty"`

	// Track is a slice of tracks that were loaded.
	Tracks []Track `json:"tracks,omitempty"`

	// Playlist contains information about the loaded playlist if one should be
	// loaded.
	Playlist *Playlist `json:"playlistInfo,omitempty"`
}
