package message

// ConfigureResuming is the message type of opcode.ConfigureResuming messages.
type ConfigureResuming struct {
	Outgoing

	// Key is the string to send when resuming a session. If the key is not
	// present, session resumption will be disabled server-side.
	Key *string `json:"key"`

	// Timeout is the number of seconds after disconnecting before the session
	// is automatically closed. This is useful for avoiding accidental leaks. It
	// defaults to 60 seconds. If the Key should not be set, this value is
	// redundant.
	Timeout uint `json:"timeout"`
}
