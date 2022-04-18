package message

// Destroy is the message type of opcode.Destroy messages.
type Destroy struct {
	Outgoing
	Guild
}
