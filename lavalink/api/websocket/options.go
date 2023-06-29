package websocket

import (
	"github.com/lukasl-dev/waterlink/v3/discord"
	"net/http"
)

// Options contains a set of configurable values to connect to a Lavalink node.
type Options struct {
	// Authorization is the password set in the configuration of the Lavalink
	// node to connect to.
	Authorization string `json:"authorization,omitempty"`

	// UserID is the user ID of the Discord bot.
	UserID discord.Snowflake `json:"userId,omitempty"`

	// ResumeKey is the configured key to resume the session with. This is
	// optional.
	ResumeKey string `json:"resumeKey,omitempty"`
}

// http returns the header as http.Header.
func (opt Options) http() http.Header {
	head := make(http.Header)
	head.Set("Authorization", opt.Authorization)
	head.Set("User-Id", opt.UserID.String())
	head.Set("Client-Name", "waterlink/v3")
	if opt.ResumeKey != "" {
		head.Set("Resume-Key", opt.ResumeKey)
	}
	return head
}
