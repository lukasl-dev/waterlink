package event

import "github.com/gompus/snowflake"

type PlayerUpdate struct {
	GuildID snowflake.Snowflake `json:"guildId,omitempty"`
	State   PlayerUpdateState   `json:"state,omitempty"`
}

type PlayerUpdateState struct {
	Time      uint `json:"time,omitempty"`
	Position  uint `json:"position,omitempty"`
	Connected bool `json:"connected,omitempty"`
}
