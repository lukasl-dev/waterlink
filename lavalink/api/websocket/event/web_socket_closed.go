package event

import "github.com/lukasl-dev/waterlink/v3/discord"

// WebSocketClosed is emitted when an audio WebSocket (to Discord) is closed.
// This can happen for various reasons (normal and abnormal), e.g. when using an
// expired voice server update. 4xx codes are usually bad.
//
// See:
// https://discordapp.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-close-event-codes
type WebSocketClosed struct {
	// Code is the Discord close event code.
	Code discord.VoiceCloseEvent `json:"code,omitempty"`

	// Reason is the reason the WebSocket was closed.
	Reason string `json:"reason,omitempty"`

	// ByRemote indicates whether the WebSocket was closed by Discord.
	ByRemote bool `json:"byRemote,omitempty"`
}
