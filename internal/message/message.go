package message

import (
	"github.com/lukasl-dev/waterlink/v2/internal/message/opcode"
)

// Incoming represents the basic attributes of an incoming message.
type Incoming struct {
	// Op determines the type of incoming messages.
	Op opcode.Incoming `json:"op"`
}

// Outgoing represents the basic attributes of an outgoing message.
type Outgoing struct {
	// Op determines the type of outgoing messages.
	Op opcode.Outgoing `json:"op"`
}

// Guild holds information of the guild for guild-scoped messages.
type Guild struct {
	// GuildID is ID of the guild from which the message was sent from/to.
	GuildID string `json:"guildId"`
}
