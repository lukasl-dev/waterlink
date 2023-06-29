package message

import (
	"github.com/lukasl-dev/waterlink/v3/discord"
	"github.com/lukasl-dev/waterlink/v3/lavalink/api/websocket/event"
)

type Event struct {
	Type    event.Type        `json:"type,omitempty"`
	GuildID discord.Snowflake `json:"guildId,omitempty"`
}
