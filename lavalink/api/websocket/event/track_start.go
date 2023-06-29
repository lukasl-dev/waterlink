package event

import (
	"github.com/lukasl-dev/waterlink/v3/discord"
	"github.com/lukasl-dev/waterlink/v3/lavalink/player"
)

// TrackStart is emitted when a track starts playing.
type TrackStart struct {
	// GuildID the guild ID of the player.
	GuildID discord.Snowflake `json:"guildId,omitempty"`

	// EncodedTrack is the base64 encoded track that starts playing.
	EncodedTrack player.EncodedTrack `json:"encodedTrack,omitempty"`
}
