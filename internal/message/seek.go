package message

// Seek is the message type of opcode.Seek messages.
type Seek struct {
	Outgoing
	Guild

	// Position is the timestamp to seek to in milliseconds.
	Position uint `json:"position"`
}
