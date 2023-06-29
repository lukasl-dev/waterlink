package lavalink

type Version struct {
	// Semver is the full version string of the Lavalink server.
	Semver VersionString `json:"semver,omitempty"`

	// Major is the major version of the Lavalink server.
	Major uint `json:"major,omitempty"`

	// Minor is the minor version of the Lavalink server.
	Minor uint `json:"minor,omitempty"`

	// Patch is the patch version of the Lavalink server.
	Patch uint `json:"patch,omitempty"`

	// PreRelease is the pre-release version according to Semver as a . separated
	// list of identifiers.
	PreRelease string `json:"preRelease,omitempty"`
}
