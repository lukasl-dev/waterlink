package waterlink

import (
	"fmt"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/waterlink/internal"
	"net/http"
)

// Credentials is a struct that holds the necessary information to communicate
// with the server.
type Credentials struct {
	// Authorization is the password matching the server configuration.
	Authorization string `json:"authorization,omitempty"`

	// UserID is the user ID of the bot to play music with.
	UserID snowflake.Snowflake `json:"userId,omitempty"`
}

// header creates a http.Header from creds and returns it.
func (creds Credentials) header() http.Header {
	h := make(http.Header)
	h.Set("Authorization", creds.Authorization)
	h.Set("User-Id", creds.UserID.String())
	h.Set("Client-Name", fmt.Sprintf("%s/%s", internal.Name, internal.Version))
	return h
}
