package waterlink

import (
	"fmt"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/waterlink/internal/message"
	"github.com/lukasl-dev/waterlink/internal/message/opcode"
	"github.com/lukasl-dev/waterlink/internal/pkgerror"
	"github.com/lukasl-dev/waterlink/track"
)

// Guild is a struct that is used to send guild-scoped messages via the
// Connection.
type Guild struct {
	// w is the json writer to write the message payloads to.
	w jsonWriter

	// id is the id of the guild to which this scope belongs to.
	id snowflake.Snowflake
}

// UpdateVoice provides an intercepted voice server update event to the server.
// This causes the server to connect to a voice channel.
func (g Guild) UpdateVoice(session string, token, endpoint string) error {
	switch {
	case session == "":
		return g.newErr("update voice", "session must be present (not empty)")
	case token == "":
		return g.newErr("update voice", "token must be present (not empty)")
	case endpoint == "":
		return g.newErr("update voice", "endpoint must be present (not empty)")
	}

	return g.wrapErr("voice update", g.w.WriteJSON(message.VoiceUpdate{
		Outgoing: message.Outgoing{Op: opcode.VoiceUpdate},
		Guild:    g.guild(),
		Session:  session,
		Event: message.VoiceUpdateEvent{
			GuildID:  g.id.String(),
			Token:    token,
			Endpoint: endpoint,
		},
	}))
}

// Play plays the preloaded audio track whose id is given via the guild's audio
// player. If no params should be given, the defaultPlayParams are used.
func (g Guild) Play(trackID string, params ...PlayParams) error {
	switch {
	case trackID == "":
		return g.newErr("play", "trackID must not be empty")
	case len(params) == 0:
		params = []PlayParams{defaultPlayParams}
	case len(params) > 1:
		return g.newErr("play", "too many params")
	}

	p := params[0]
	return g.wrapErr("play", g.w.WriteJSON(message.Play{
		Outgoing:  message.Outgoing{Op: opcode.Play},
		Guild:     g.guild(),
		Track:     trackID,
		StartTime: p.startTime(),
		EndTime:   p.endTime(),
		Volume:    p.volume(),
		NoReplace: p.Pause,
		Pause:     p.Pause,
	}))
}

// PlayTrack plays the given audio track. If no params should be given, the
// defaultPlayParams are used.
func (g Guild) PlayTrack(tr track.Track, params ...PlayParams) error {
	return g.Play(tr.ID, params...)
}

// Stop stops the audio playback of the guild's audio player.
func (g Guild) Stop() error {
	return g.wrapErr("stop", g.w.WriteJSON(message.Stop{
		Outgoing: message.Outgoing{Op: opcode.Stop},
		Guild:    g.guild(),
	}))
}

// Pause pauses the audio playback of the guild's audio player.
func (g Guild) Pause() error {
	return g.wrapErr("pause", g.w.WriteJSON(message.Pause{
		Outgoing: message.Outgoing{Op: opcode.Pause},
		Guild:    g.guild(),
	}))
}

// Seek seeks the current playing audio track to a specific position in
// milliseconds.
func (g Guild) Seek(position uint) error {
	return g.wrapErr("seek", g.w.WriteJSON(message.Seek{
		Outgoing: message.Outgoing{Op: opcode.Seek},
		Guild:    g.guild(),
		Position: position,
	}))
}

// Volume updates the volume of the guild's audio player. The value must be
// between 0 and 1000. Defaults to 100.
func (g Guild) Volume(volume uint16) error {
	if volume > 1000 {
		return g.newErr("volume", "volume must be between 0 and 1000")
	}

	return g.wrapErr("volume", g.w.WriteJSON(message.Volume{
		Outgoing: message.Outgoing{Op: opcode.Volume},
		Guild:    g.guild(),
		Volume:   volume,
	}))
}

// newErr creates a new error with action as prefix.
func (g Guild) newErr(action, msg string) error {
	return pkgerror.New(fmt.Sprintf("connection: guild: %s: %s: %s", action, g.id, msg))
}

// wrapErr wraps err and adds action as prefix to the error message.
func (g Guild) wrapErr(action string, err error) error {
	return pkgerror.Wrap(fmt.Sprintf("connection: guild: %s: %s:", action, g.id), err)
}

// guild creates a message.Guild from g.
func (g Guild) guild() message.Guild {
	return message.Guild{GuildID: g.id.String()}
}
