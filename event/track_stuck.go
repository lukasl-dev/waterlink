package event

import "github.com/gompus/snowflake"

type TrackStuck struct {
	GuildID     snowflake.Snowflake `json:"guildId,omitempty"`
	TrackID     string              `json:"trackId,omitempty"`
	ThresholdMS uint                `json:"thresholdMs,omitempty"`
}
