package lavalink

import (
	"encoding/json"
	"time"
)

type Git struct {
	// Branch is the Git branch based on which the Lavalink server was built.
	Branch string `json:"branch,omitempty"`

	// Commit is the Git commit hash based on which the Lavalink server was
	// built.
	Commit string `json:"commit,omitempty"`

	// CommitTime is the timestamp when the Git commit was made.
	CommitTime time.Time `json:"commitTime,omitempty"`
}

var (
	_ json.Marshaler   = (*Git)(nil)
	_ json.Unmarshaler = (*Git)(nil)
)

type git struct {
	Branch     string `json:"branch,omitempty"`
	Commit     string `json:"commit,omitempty"`
	CommitTime int    `json:"commitTime,omitempty"`
}

func (g *Git) MarshalJSON() ([]byte, error) {
	return json.Marshal(git{
		Branch:     g.Branch,
		Commit:     g.Commit,
		CommitTime: int(g.CommitTime.UnixMilli()),
	})
}

func (g *Git) UnmarshalJSON(bytes []byte) error {
	var u git
	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	g.Branch = u.Branch
	g.Commit = u.Commit
	g.CommitTime = time.UnixMilli(int64(u.CommitTime))

	return nil
}
