package track

// LoadType determines the behaviour of LoadResult.
type LoadType string

const (
	// TracksLoaded is used when the tracks have been loaded.
	TracksLoaded LoadType = "TRACKS_LOADED"

	// PlaylistLoaded is used when the playlist (and its tracks) has been loaded.
	PlaylistLoaded LoadType = "PLAYLIST_LOADED"

	// SearchResult is used when the search result has been resolved.
	SearchResult LoadType = "SEARCH_RESULT"

	// NoMatches is used when no matches were found.
	NoMatches LoadType = "NO_MATCHES"

	// LoadFailed is used when the load failed.
	LoadFailed LoadType = "LOAD_FAILED"
)
