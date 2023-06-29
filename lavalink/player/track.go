package player

import (
	"encoding/json"
	"github.com/lukasl-dev/waterlink/v3/lavalink"
	"net/url"
	"time"
)

type Track struct {
	// Encoded is the base64-encoded track.
	Encoded EncodedTrack `json:"encoded,omitempty"`

	// Info holds information about the track.
	Info TrackInfo `json:"info,omitempty"`
}

type TrackInfo struct {
	// Identifier is the track identifier.
	Identifier Query `json:"identifier,omitempty"`

	// Seekable indicates if the track is seekable.
	Seekable bool `json:"isSeekable,omitempty"`

	// Author is the author of the track.
	Author string `json:"author,omitempty"`

	// Length is the length of the track.
	Length time.Duration `json:"length,omitempty"`

	// Stream indicates if the track is a stream.
	Stream bool `json:"isStream,omitempty"`

	// Position is the position of the track.
	Position time.Duration `json:"position,omitempty"`

	// Title is the title of the track.
	Title string `json:"title,omitempty"`

	// URI is the URI of the track.
	URI *url.URL `json:"uri,omitempty"`

	// SourceName is the name of the source of the track.
	SourceName lavalink.SourceManager `json:"sourceName,omitempty"`
}

var (
	_ json.Marshaler   = (*Track)(nil)
	_ json.Unmarshaler = (*Track)(nil)
)

type track struct {
	Encoded string    `json:"encoded,omitempty"`
	Info    trackInfo `json:"info,omitempty"`
}

type trackInfo struct {
	Identifier Query                  `json:"identifier,omitempty"`
	IsSeekable bool                   `json:"isSeekable,omitempty"`
	Author     string                 `json:"author,omitempty"`
	Length     uint                   `json:"length,omitempty"`
	IsStream   bool                   `json:"isStream,omitempty"`
	Position   uint                   `json:"position,omitempty"`
	Title      string                 `json:"title,omitempty"`
	URI        string                 `json:"uri,omitempty"`
	SourceName lavalink.SourceManager `json:"sourceName,omitempty"`
}

func (t *Track) MarshalJSON() ([]byte, error) {
	return json.Marshal(track{
		Encoded: string(t.Encoded),
		Info: trackInfo{
			Identifier: t.Info.Identifier,
			IsSeekable: t.Info.Seekable,
			Author:     t.Info.Author,
			Length:     uint(t.Info.Length.Milliseconds()),
			IsStream:   t.Info.Stream,
			Position:   uint(t.Info.Position.Milliseconds()),
			Title:      t.Info.Title,
			URI:        t.Info.URI.String(),
			SourceName: t.Info.SourceName,
		},
	})
}

func (t *Track) UnmarshalJSON(bytes []byte) error {
	var u track
	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	uri, err := url.ParseRequestURI(u.Info.URI)
	if err != nil {
		return err
	}

	t.Encoded = EncodedTrack(u.Encoded)
	t.Info = TrackInfo{
		Identifier: u.Info.Identifier,
		Seekable:   u.Info.IsSeekable,
		Author:     u.Info.Author,
		Length:     time.Duration(u.Info.Length) * time.Millisecond,
		Stream:     u.Info.IsStream,
		Position:   time.Duration(u.Info.Position) * time.Millisecond,
		Title:      u.Info.Title,
		URI:        uri,
		SourceName: u.Info.SourceName,
	}

	return nil
}
