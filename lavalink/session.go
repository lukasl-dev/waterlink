package lavalink

import (
	"encoding/json"
	"time"
)

type Session struct {
	// ResumeKey is the key to use to resume the session. If empty, the session
	// cannot be resumed.
	ResumeKey string `json:"resumeKey,omitempty"`

	// Timeout is the duration to keep the session alive for after a disconnect.
	// Defaults to 60s.
	Timeout time.Duration `json:"timeout,omitempty"`
}

var (
	_ json.Marshaler   = (*Session)(nil)
	_ json.Unmarshaler = (*Session)(nil)
)

type session struct {
	ResumeKey string `json:"resumeKey,omitempty"`
	Timeout   uint   `json:"timeout,omitempty"`
}

func (s *Session) MarshalJSON() ([]byte, error) {
	return json.Marshal(session{
		ResumeKey: s.ResumeKey,
		Timeout:   uint(s.Timeout.Seconds()),
	})
}

func (s *Session) UnmarshalJSON(bytes []byte) error {
	var u session
	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	s.ResumeKey = u.ResumeKey
	s.Timeout = time.Duration(u.Timeout) * time.Second

	return nil
}
