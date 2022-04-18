package event

import "github.com/gompus/snowflake"

type TrackStart struct {
	GuildID snowflake.Snowflake `json:"guildId,omitempty"`
	TrackID string              `json:"track,omitempty"`
}
