package waterlink

type LoadType string

const (
	LoadTypeUnknown        LoadType = ""
	LoadTypeTrackLoaded             = "TRACK_LOADED "
	LoadTypePlaylistLoaded          = "PLAYLIST_LOADED"
	LoadTypeSearchResult            = "SEARCH_RESULT "
	LoadTypeNoMatches               = "NO_MATCHES"
	LoadTypeLoadFailed              = "LOAD_FAILED"
)
