package lavalink

type Plugin struct {
	// Name is the name of the plugin.
	Name string `json:"name,omitempty"`

	// Version is the version of the plugin.
	Version string `json:"version,omitempty"`
}
