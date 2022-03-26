package opcode

// Incoming determines the type of incoming messages.
type Incoming string

const (
	// PlayerUpdate contains information about a guild's audio player.
	PlayerUpdate Incoming = "playerUpdate"

	// Stats contains information about the server instance.
	Stats Incoming = "stats"

	// Event contains a server-side occurred event that is delegated to the
	// client.
	Event Incoming = "event"
)
