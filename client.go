package waterlink

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/equalizer"
	"github.com/lukasl-dev/waterlink/play"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var (
	ErrLoadFailed = errors.New("player loading failed: %w")
	ErrNoMatches  = errors.New("no matches found")
)

type Client struct {
	wsHost      string
	httpHost    string
	conn        *websocket.Conn
	password    string
	userID      string
	totalShards int
}

// newPlay returns a new client with passed options.
func New(options ...Option) (*Client, error) {
	c := &Client{
		httpHost: "http://localhost:8080",
		wsHost:   "ws://localhost:8080",
	}

	// apply options
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) authorize(header http.Header) {
	header.Set("Authorization", c.password)
}

// Open opens the connection to the websocket server.
func (c *Client) Open() error {
	header := http.Header{}
	c.authorize(header)
	header.Set("Num-Shards", strconv.FormatInt(int64(c.totalShards), 10))
	header.Set("User-Id", c.userID)

	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(c.wsHost, header)
	if err != nil {
		return err
	}

	return nil
}

// send sends v as json to the websocket server.
func (c *Client) send(v interface{}) error {
	return c.conn.WriteJSON(v)
}

// VoiceUpdate provides a voice server update.
func (c *Client) VoiceUpdate(guildID string, sessionID string, event VoiceServerUpdate) error {
	type body struct {
		OP        op                `json:"op"`
		GuildID   string            `json:"guildId"`
		SessionID string            `json:"sessionId"`
		Event     VoiceServerUpdate `json:"event"`
	}

	return c.send(body{
		OP:        opVoiceUpdate,
		GuildID:   guildID,
		SessionID: sessionID,
		Event:     event,
	})
}

// Play plays a specific track on a specific guild.
func (c *Client) Play(guildID string, trackID string, options ...play.Option) error {
	p := play.New(guildID, trackID)

	for _, option := range options {
		if err := option(p); err != nil {
			return err
		}
	}

	return c.send(p)
}

// Play plays a specific track on a specific guild.
func (c *Client) PlayTrack(guildID string, track Track, options ...play.Option) error {
	return c.Play(guildID, track.ID, options...)
}

// Stop stops the player.
func (c *Client) Stop(guildID string) error {
	type body struct {
		OP      op     `json:"op"`
		GuildID string `json:"guildId"`
	}

	return c.send(body{
		OP:      opStop,
		GuildID: guildID,
	})
}

// Pause pauses the playback.
func (c *Client) Pause(guildID string, pause bool) error {
	type body struct {
		OP      op     `json:"op"`
		GuildID string `json:"guildId"`
		Pause   bool   `json:"pause"`
	}

	return c.send(body{
		OP:      opPause,
		GuildID: guildID,
		Pause:   pause,
	})
}

// Seek seeks a player.
func (c *Client) Seek(guildID string, position int) error {
	type body struct {
		OP       op     `json:"op"`
		GuildID  string `json:"guildId"`
		Position int    `json:"position"`
	}

	return c.send(body{
		OP:       opSeek,
		GuildID:  guildID,
		Position: position,
	})
}

// Volume sets the player's volume.
func (c *Client) Volume(guildID string, volume int) error {
	type body struct {
		OP      op     `json:"op"`
		GuildID string `json:"guildId"`
		Volume  int    `json:"volume"`
	}

	return c.send(body{
		OP:      opVolume,
		GuildID: guildID,
		Volume:  volume,
	})
}

// Equalizer uses the player's equalizer.
//
// There are 15 bands (0-14) that can be changed.
// Gain is the multiplier for the given band.
// The default value is 0. Valid values range from -0.25 to 1.0, where -0.25 means the given band is completely muted,
// and 0.25 means it is doubled.
// Modifying the gain could also change the volume of the output.
func (c *Client) Equalizer(guildID string, bands ...equalizer.Band) error {
	type body struct {
		OP      op               `json:"op"`
		GuildID string           `json:"guildId"`
		Bands   []equalizer.Band `json:"bands"`
	}

	return c.send(body{
		OP:      opEqualizer,
		GuildID: guildID,
		Bands:   bands,
	})
}

// Destroy destroys a player.
//
// Tell the server to potentially disconnect from the voice server and potentially remove the player with all its data.
// This is useful if you want to move to a new node for a voice connection.
// Calling this op does not affect voice state, and you can send the same VOICE_SERVER_UPDATE to a new node.
func (c *Client) Destroy(guildID string) error {
	type body struct {
		OP      op     `json:"op"`
		GuildID string `json:"guildId"`
	}

	return c.send(body{
		OP:      opDestroy,
		GuildID: guildID,
	})
}

func (c *Client) LoadTracks(identifier string) (LoadType, PlaylistInfo, []Track, error) {
	type body struct {
		LoadType     LoadType     `json:"loadType"`
		PlaylistInfo PlaylistInfo `json:"playlistInfo"`
		Tracks       []Track      `json:"tracks"`
		Exception    struct {
			Message  string `json:"message"`
			Severity string `json:"severity"`
		} `json:"exception"`
	}

	host, _ := url.Parse(c.httpHost)
	host.Path = "/loadtracks"
	query := host.Query()
	query.Set("identifier", identifier)
	host.RawQuery = query.Encode()

	request, err := http.NewRequest("GET", host.String(), nil)
	if err != nil {
		return LoadTypeUnknown, PlaylistInfo{}, nil, err
	}
	c.authorize(request.Header)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return LoadTypeUnknown, PlaylistInfo{}, nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return LoadTypeUnknown, PlaylistInfo{}, nil, err
	}

	var marshal *body
	if err := json.Unmarshal(responseBody, &marshal); err != nil {
		return LoadTypeUnknown, PlaylistInfo{}, nil, err
	}

	switch marshal.LoadType {
	case LoadTypeNoMatches:
		return LoadTypeNoMatches, PlaylistInfo{}, nil, ErrNoMatches
	case LoadTypeLoadFailed:
		return LoadTypeLoadFailed, PlaylistInfo{}, nil, fmt.Errorf(
			"%s: %s",
			marshal.Exception.Severity, marshal.Exception.Message,
		)
	}

	return marshal.LoadType, marshal.PlaylistInfo, marshal.Tracks, nil
}
