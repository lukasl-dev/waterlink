package message

// VoiceUpdate is the message type of opcode.VoiceUpdate messages.
type VoiceUpdate struct {
	Outgoing
	Guild

	Session string `json:"session"`

	// Event is the intercepted voice server update event.
	Event VoiceUpdateEvent `json:"event"`
}

type VoiceUpdateEvent struct {
	GuildID  string `json:"guild_id"`
	Token    string `json:"token"`
	Endpoint string `json:"endpoint"`
}
