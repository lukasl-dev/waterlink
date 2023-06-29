package player

import "github.com/lukasl-dev/waterlink/v3/lavalink"

type TrackLoadResult struct {
	LoadType     TrackLoadResultType `json:"loadType,omitempty"`
	PlaylistInfo PlaylistInfo        `json:"playlistInfo,omitempty"`
	Tracks       []Track             `json:"tracks,omitempty"`
	Exception    *lavalink.Exception `json:"exception,omitempty"`
}

type TrackLoadResultType string

const (
	TrackLoadResultTypeTrackLoaded    TrackLoadResultType = "TRACK_LOADED"
	TrackLoadResultTypePlaylistLoaded TrackLoadResultType = "PLAYLIST_LOADED"
	TrackLoadResultTypeSearchResult   TrackLoadResultType = "SEARCH_RESULT"
	TrackLoadResultTypeNoMatches      TrackLoadResultType = "NO_MATCHES"
	TrackLoadResultTypeLoadFailed     TrackLoadResultType = "LOAD_FAILED"
)
