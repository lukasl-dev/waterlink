package event

import "github.com/gompus/snowflake"

type TrackException struct {
	GuildID snowflake.Snowflake `json:"guildId,omitempty"`
	TrackID string              `json:"track,omitempty"`
	Error   string              `json:"error,omitempty"`
}
