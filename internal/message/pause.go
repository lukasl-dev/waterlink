package message

// Pause is the message type of opcode.Pause messages.
type Pause struct {
	Outgoing
	Guild

	// Pause indicates whether the playback should be paused or resumed.
	Pause bool `json:"pause"`
}
