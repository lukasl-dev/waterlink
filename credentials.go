package waterlink

import (
	"fmt"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/waterlink/v2/internal"
	"net/http"
)

// Credentials is a struct that holds the necessary information to communicate
// with the server.
type Credentials struct {
	// Authorization is the password matching the server configuration. This
	// is required for the initial handshake.
	Authorization string `json:"authorization,omitempty"`

	// UserID is the user ID of the bot to play music with. This is required
	// to play audio.
	UserID snowflake.Snowflake `json:"userId,omitempty"`

	// ResumeKey is the key of the session you want to resume. If it is empty,
	// a new session will be created.
	ResumeKey string `json:"resumeKey,omitempty"`
}

// header creates a http.Header from creds and returns it.
func (creds Credentials) header() http.Header {
	h := make(http.Header)
	h.Set("Client-Name", fmt.Sprintf("%s/%s", internal.Name, internal.Version))
	h.Set("Authorization", creds.Authorization)
	h.Set("User-Id", creds.UserID.String())
	if creds.ResumeKey != "" {
		h.Set("Resume-Key", creds.ResumeKey)
	}
	return h
}
