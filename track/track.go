package track

import (
	"github.com/lukasl-dev/waterlink/v2/track/query"
)

// Track represents a preloaded track.
type Track struct {
	// ID is the internal ID of the track. It is used to play the track via
	// a guild's audio track.
	ID string `json:"track,omitempty"`

	// Info is the track's information.
	Info Info `json:"info,omitempty"`
}

// Info contains information about a track.
type Info struct {
	// Query is the query that was used to resolve this track.
	Query query.Query `json:"identifier,omitempty"`

	// Seekable indicates whether the track can be seeked to a specific time.
	Seekable bool `json:"isSeekable,omitempty"`

	// Author is the author of the track. For example: "RickAstleyVEVO"
	Author string `json:"author,omitempty"`

	// Length is the length of the track in milliseconds.
	Length uint `json:"length,omitempty"`

	// Stream indicates whether the audio source is a stream.
	Stream bool `json:"isStream,omitempty"`

	// Title is the title of the track. For example: "Rick Astley - Never Gonna
	// Give You Up".
	Title string `json:"title,omitempty"`

	// URI is the URL to the audio source/stream of the track.
	URI string `json:"uri,omitempty"`

	// SourceName is the name of the source/platform from which the track was
	// retrieved. For example: "youtube".
	SourceName string `json:"sourceName,omitempty"`
}
