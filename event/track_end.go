package event

import "github.com/gompus/snowflake"

type TrackEnd struct {
	GuildID snowflake.Snowflake `json:"guildId,omitempty"`
	TrackID string              `json:"track,omitempty"`
}
