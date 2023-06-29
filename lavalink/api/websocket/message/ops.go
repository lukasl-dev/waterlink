package message

// Op represents an opcode for a websocket payload. It is used to distinguish
// between different messages.
type Op string

const (
	// OpReady is emitted you successfully connect to the Lavalink node.
	OpReady Op = "ready"

	// OpPlayerUpdate is emitted every x seconds with the latest player state.
	OpPlayerUpdate Op = "playerUpdate"

	// OpStats is emitted when a player or voice event is emitted.
	OpStats Op = "stats"

	// OpEvent is emitted when an event is emitted.
	OpEvent Op = "event"
)
