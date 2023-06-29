package player

import "github.com/lukasl-dev/waterlink/v3/discord"

type Player struct {
	// GuildID of the guild ID of the player.
	GuildID discord.Snowflake `json:"guildId,omitempty"`

	// Track is the currently playing track.
	Track *Track `json:"track,omitempty"`

	// Volume is the volume in percentage of the player. Range is 0-1000.
	Volume uint `json:"volume,omitempty"`

	// Paused is whether the player is paused.
	Paused bool `json:"paused,omitempty"`

	// Voice is the voice state of the player.
	Voice VoiceState `json:"voice,omitempty"`

	// Filters is the filters used by the player.
	Filters Filters `json:"filters,omitempty"`
}
