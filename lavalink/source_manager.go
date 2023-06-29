package lavalink

// SourceManager represents a possible source manager of a Lavalink node.
type SourceManager string

const (
	SourceManagerYouTube    SourceManager = "youtube"
	SourceManagerSoundCloud SourceManager = "soundcloud"
	SourceManagerBandCamp   SourceManager = "bandcamp"
	SourceManagerTwitch     SourceManager = "twitch"
	SourceManagerVimeo      SourceManager = "vimeo"
	SourceManagerBeamPro    SourceManager = "beam.pro"
	SourceManagerHTTP       SourceManager = "http"
)
