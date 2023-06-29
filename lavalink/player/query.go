package player

// Query represents an identifier to load a track from.
type Query string

// RawQuery returns a Query that will load the given identifier directly.
func RawQuery(query string) Query {
	return Query(query)
}

// YouTubeQuery returns a Query that will load the given identifier as a
// YouTube search query.
func YouTubeQuery(query string) Query {
	return Query("ytsearch:" + query)
}

// SoundCloudQuery returns a Query that will load the given identifier as a
// SoundCloud search query.
func SoundCloudQuery(query string) Query {
	return Query("scsearch:" + query)
}
