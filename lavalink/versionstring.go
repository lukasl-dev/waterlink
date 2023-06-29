package lavalink

// VersionString is a formatted version string.
//
// Example: "3.0.1"
type VersionString string

// IsValid returns whether the version is valid.
func (v VersionString) IsValid() bool {
	return v != ""
}
