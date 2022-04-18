package message

// Volume is the message type of opcode.Volume messages.
type Volume struct {
	Outgoing
	Guild

	// Volume is the new volume of the player. The value must be between 0 and
	// 1000. Defaults to 100.
	Volume uint16 `json:"volume"`
}
