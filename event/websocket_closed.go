package event

import "github.com/gompus/snowflake"

type WebSocketClosed struct {
	GuildID snowflake.Snowflake `json:"guildId,omitempty"`
	Code    uint                `json:"code,omitempty"`
	Reason  string              `json:"reason,omitempty"`
	Remote  bool                `json:"byRemote,omitempty"`
}
