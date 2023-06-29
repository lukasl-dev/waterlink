package lavalink

import (
	"encoding/json"
	"time"
)

type Info struct {
	// Version is the version of the Lavalink server.
	Version Version `json:"version,omitempty"`

	// BuildTime is the timestamp when the Lavalink jar was built.
	BuildTime time.Time `json:"buildTime,omitempty"`

	// Git holds information about Git of the Lavalink server.
	Git Git `json:"git,omitempty"`

	// JVM is the JVM version of the Lavalink server.
	JVM string `json:"jvm,omitempty"`

	// Lavaplayer is the Lavaplayer version used by the Lavalink server.
	Lavaplayer string `json:"lavaplayer,omitempty"`

	// SourceManagers is an array of enabled source managers of the Lavalink
	// server.
	SourceManagers []SourceManager `json:"sourceManagers,omitempty"`

	// Filters is an array of enabled filters of the Lavalink server.
	Filters []string `json:"filters,omitempty"`

	// Plugins is an array of enabled plugins of the Lavalink server.
	Plugins []Plugin `json:"plugins,omitempty"`
}

var (
	_ json.Marshaler   = (*Info)(nil)
	_ json.Unmarshaler = (*Info)(nil)
)

type info struct {
	Version        Version         `json:"version,omitempty"`
	BuildTime      int             `json:"buildTime,omitempty"`
	Git            Git             `json:"git,omitempty"`
	JVM            string          `json:"jvm,omitempty"`
	Lavaplayer     string          `json:"lavaplayer,omitempty"`
	SourceManagers []SourceManager `json:"sourceManagers,omitempty"`
	Filters        []string        `json:"filters,omitempty"`
	Plugins        []Plugin        `json:"plugins,omitempty"`
}

func (i *Info) MarshalJSON() ([]byte, error) {
	return json.Marshal(info{
		Version:        i.Version,
		BuildTime:      int(i.BuildTime.UnixMilli()),
		Git:            i.Git,
		JVM:            i.JVM,
		Lavaplayer:     i.Lavaplayer,
		SourceManagers: i.SourceManagers,
		Filters:        i.Filters,
		Plugins:        i.Plugins,
	})
}

func (i *Info) UnmarshalJSON(bytes []byte) error {
	var u info
	if err := json.Unmarshal(bytes, &u); err != nil {
		return err
	}

	i.Version = u.Version
	i.BuildTime = time.UnixMilli(int64(u.BuildTime))
	i.Git = u.Git
	i.JVM = u.JVM
	i.Lavaplayer = u.Lavaplayer
	i.SourceManagers = u.SourceManagers
	i.Filters = u.Filters
	i.Plugins = u.Plugins

	return nil
}
