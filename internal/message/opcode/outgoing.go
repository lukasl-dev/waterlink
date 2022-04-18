package opcode

// Outgoing determines the type of outgoing messages.
type Outgoing string

const (
	// VoiceUpdate is used to provide an intercepted voice server update to the
	// server. This causes the server to connect a voice channel.
	VoiceUpdate Outgoing = "voiceUpdate"

	// Play is used to request the server to play a specific audio track.
	Play Outgoing = "play"

	// Stop stops the audio playback.
	Stop Outgoing = "stop"

	// Pause pauses the audio playback.
	Pause Outgoing = "pause"

	// Seek seeks the current playing audio track to a specific position in
	// milliseconds.
	Seek Outgoing = "seek"

	// Volume updates the volume of the audio playback.
	Volume Outgoing = "volume"

	// Filters applies a set of filters to the audio playback. Adding a filter
	// can have adverse effects on performance. Filters force the player to
	// decode all audio to PCM, event if the input was already in the Opus format
	// that Discord uses. his means decoding and encoding audio that would
	// normally require very little processing. This is often the case with
	// YouTube videos.
	Filters Outgoing = "filters"

	// Destroy destroys a specific audio player instance server-side.
	Destroy Outgoing = "destroy"

	// ConfigureResuming configures the session resumption behavior of the
	// server. It can enable or disable the resumption of sessions, and it can
	// define the number of seconds after which a session will be considered
	// expired server-side.
	ConfigureResuming Outgoing = "configureResuming"
)
