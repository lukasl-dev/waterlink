package discord

// VoiceCloseEvent represents a voice close event code as defined by Discord.
//
// See:
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-close-event-codes
type VoiceCloseEvent uint

const (
	VoiceCloseEventUnknownOpCode         VoiceCloseEvent = 4001
	VoiceCloseEventFailedToDecodePayload VoiceCloseEvent = 4002
	VoiceCloseEventNotAuthenticated      VoiceCloseEvent = 4003
	VoiceCloseEventAuthenticationFailed  VoiceCloseEvent = 4004
	VoiceCloseEventAlreadyAuthenticated  VoiceCloseEvent = 4005
	VoiceCloseEventSessionNoLongerValid  VoiceCloseEvent = 4006
	VoiceCloseEventSessionTimeout        VoiceCloseEvent = 4009
	VoiceCloseEventServerNotFound        VoiceCloseEvent = 4011
	VoiceCloseEventUnknownProtocol       VoiceCloseEvent = 4012
	VoiceCloseEventDisconnected          VoiceCloseEvent = 4014
	VoiceCloseEventVoiceServerCrashed    VoiceCloseEvent = 4015
	VoiceCloseEventUnknownEncryptionMode VoiceCloseEvent = 4016
)
