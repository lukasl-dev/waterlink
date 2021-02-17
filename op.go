package waterlink

type op string

//goland:noinspection ALL
var (
	// opVoiceUpdate provides an intercepted voice server update.
	// An outgoing op.
	opVoiceUpdate op = "voiceUpdate"

	// opPlay plays a Track.
	// An outgoing op.
	opPlay op = "play"

	// opStop stops a Player.
	// An outgoing op.
	opStop op = "stop"

	// opPause pauses the playback.
	// An outgoing op.
	opPause op = "pause"

	// opSeek seeks a Track.
	// An outgoing op.
	opSeek op = "seek"

	// opVolume sets the Player'name volume.
	// An outgoing op.
	opVolume op = "volume"

	// opEqualizer uses the player equalizer.
	// An outgoing op.
	opEqualizer op = "equalizer"

	// opDestroy destroys a Player.
	// An outgoing op.
	opDestroy op = "destroy"

	// opPlayerUpdate returns the position information about a Player.
	// An incoming op.
	opPlayerUpdate op = "playerUpdate"

	// opStats returns a collection of sats sent every minute.
	// An incoming op.
	opStats op = "stats"

	// opEvent is used to receive an Event emitted by de Server.
	// An incoming op.
	opEvent op = "event"
)

// String returns the op as name.
func (op op) String() string {
	return string(op)
}
