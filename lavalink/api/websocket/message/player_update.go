package message

import (
	"encoding/json"
	"github.com/lukasl-dev/waterlink/v3/discord"
	"time"
)

// PlayerUpdate is dispatched every x seconds (configurable in application.yml)
// seconds with the current state of the player.
type PlayerUpdate struct {
	// GuildID is the guild ID of the player.
	GuildID discord.Snowflake

	// State is the state of the player.
	State PlayerUpdateState
}

type PlayerUpdateState struct {
	// Time is the unix timestamp.
	Time time.Time `json:"time,omitempty"`

	// Position is the position of the track.
	Position time.Duration `json:"position,omitempty"`

	// Connected is whether Lavalink is connected to the voice gateway.
	Connected bool `json:"connected,omitempty"`

	// The ping of the node to the discord voice server in milliseconds. If not
	// connected, this is -1.
	Ping int `json:"ping,omitempty"`
}

var (
	_ json.Marshaler   = (*PlayerUpdateState)(nil)
	_ json.Unmarshaler = (*PlayerUpdateState)(nil)
)

type playerUpdateState struct {
	Time      uint `json:"time"`
	Position  uint `json:"position"`
	Connected bool `json:"connected"`
	Ping      int  `json:"ping"`
}

func (pus *PlayerUpdateState) MarshalJSON() ([]byte, error) {
	return json.Marshal(playerUpdateState{
		Time:      uint(pus.Time.UnixMilli()),
		Position:  uint(pus.Position.Milliseconds()),
		Connected: pus.Connected,
		Ping:      pus.Ping,
	})
}

func (pus *PlayerUpdateState) UnmarshalJSON(b []byte) error {
	var u playerUpdateState
	if err := json.Unmarshal(b, &u); err != nil {
		return err
	}

	pus.Time = time.UnixMilli(int64(u.Time))
	pus.Position = time.Duration(u.Position) * time.Millisecond
	pus.Connected = u.Connected
	pus.Ping = u.Ping

	return nil
}
