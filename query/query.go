package query

import "fmt"

// Query represents a string used to load audio tracks from the server.
type Query string

// YouTube returns a query that searches for YouTube videos matching the given
// search query.
func YouTube(search string) Query {
	return Query(fmt.Sprintf("ytsearch: %s", search))
}

// SoundCloud returns a query that searches for SoundCloud tracks matching the
// given search query.
func SoundCloud(search string) Query {
	return Query(fmt.Sprintf("scsearch: %s", search))
}
