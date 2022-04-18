package message

// Stop is the message type of opcode.Stop messages.
type Stop struct {
	Outgoing
	Guild
}
