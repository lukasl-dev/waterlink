package rest

import (
	"fmt"
	"github.com/lukasl-dev/waterlink/v3/discord"
	"github.com/lukasl-dev/waterlink/v3/internal"
	"github.com/lukasl-dev/waterlink/v3/lavalink"
	"github.com/lukasl-dev/waterlink/v3/lavalink/player"
	"github.com/lukasl-dev/waterlink/v3/lavalink/routeplanner"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"net/http"
)

type Client struct {
	// client is the client used to dispatch HTTP requests.
	client *gentleman.Client
}

var (
	// ErrClientInvalidAddress occurs when an invalid address is passed to
	// Open().
	//
	// Valid address:
	// ws://localhost:2333
	//
	// Invalid address:
	// localhost:2333
	ErrClientInvalidAddress = fmt.Errorf("waterlink: rest: address is invalid")
)

// NewClient creates a new REST Client which dispatches requests to the given
// address with the given header values. It returns an error if the connection
// could not be established.
//
// Example address: http://localhost:2333
func NewClient(addr string, opt Options) (*Client, error) {
	if !internal.IsURL(addr) {
		return nil, ErrClientInvalidAddress
	}

	client := gentleman.New().
		URL(addr).
		SetHeader("Authorization", opt.Authorization)

	return &Client{client: client}, nil
}

// Info returns information about the node. It dispatches a GET request to
// /v3/info.
func (cl *Client) Info() (*lavalink.Info, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path("/v3/info")

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: node api: failed to get info: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: node api: failed to get info: unexpected status code %d received", resp.StatusCode)
	}

	var info lavalink.Info
	if err := resp.JSON(&info); err != nil {
		return nil, fmt.Errorf("waterlink: rest: node api: failed to get info: %w", err)
	}

	return &info, nil
}

// Stats returns the stats of the node. It dispatches a GET request to /v3/stats.
// Note that the lavalink.FrameStats field in lavalink.Stats is always nil.
func (cl *Client) Stats() (*lavalink.Stats, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path("/v3/stats")

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: node api: failed to get info: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: node api: failed to get info: unexpected status code %d received", resp.StatusCode)
	}

	var stats lavalink.Stats
	if err := resp.JSON(&stats); err != nil {
		return nil, fmt.Errorf("waterlink: rest: node api: failed to get info: %w", err)
	}

	return &stats, nil
}

func (cl *Client) Version() (lavalink.VersionString, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path("/version")

	resp, err := req.Send()
	if err != nil {
		return "", fmt.Errorf("waterlink: rest: node api: failed to get version: %w", err)
	}

	if !resp.Ok {
		return "", fmt.Errorf("waterlink: rest: node api: failed to get version: unexpected status code %d received", resp.StatusCode)
	}

	version := lavalink.VersionString(resp.String())
	return version, nil
}

// Players returns all players in the current session. It dispatches a GET request
// to /sessions/{sessionID}/players, whereby {sessionID} is equal to the ID
// of the session to query for players.
func (cl *Client) Players(sessionID string) ([]player.Player, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path(fmt.Sprintf("/v3/sessions/%s/players", sessionID))

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: player api: failed to get all players: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: player api: failed to get all players: unexpected status code %d received", resp.StatusCode)
	}

	var players []player.Player
	if err := resp.JSON(&players); err != nil {
		return nil, fmt.Errorf("waterlink: rest: player api: failed to get all players: %w", err)
	}

	return players, nil
}

// Player returns the player for the specified guild. It dispatches a GET request
// to /sessions/{sessionID}/players/{guildID}, whereby {sessionID} is equal to
// the ID of the session to query for the player and {guildID} is the ID of the
// guild to query the player for.
func (cl *Client) Player(sessionID string, guildID discord.Snowflake) (*player.Player, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path(fmt.Sprintf("/v3/sessions/%s/players/%s", sessionID, guildID))

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: player api: failed to get player: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: player api: failed to get player: unexpected status code %d received", resp.StatusCode)
	}

	var play player.Player
	if err := resp.JSON(&play); err != nil {
		return nil, fmt.Errorf("waterlink: rest: player api: failed to get player: %w", err)
	}

	return &play, nil
}

//func (api *Player) Update(guildID discord.Snowflake, update player.Update, noReplace bool) {
//	//path := fmt.Sprintf("/sessions/%s/players/%s?noReplace=%s", api.sessionID, guildID, strconv.FormatBool(noReplace))
//
//	// TODO: implement
//	panic("not implemented")
//}

// DestroyPlayer destroys the player for the specified guild. It dispatches a DELETE
// request to /sessions/{sessionID}/players/{guildID}, whereby {sessionID} is
// always equal to the sessionID defined in the Player API and {guildID} is the
// ID of the guild to destroy the player for.
func (cl *Client) DestroyPlayer(sessionID string, guildID discord.Snowflake) error {
	req := cl.client.Request().
		Method(http.MethodDelete).
		Path(fmt.Sprintf("/v3/sessions/%s/players/%s", sessionID, guildID))

	resp, err := req.Send()
	if err != nil {
		return fmt.Errorf("waterlink: rest: player api: failed to destroy player: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("waterlink: rest: player api: failed to destroy player: unexpected status code %d received", resp.StatusCode)
	}

	return nil
}

// RoutePlannerStatus returns the status of the route planner. It dispatches a
// GET request to /v3/routeplanner/status.
func (cl *Client) RoutePlannerStatus() (*routeplanner.Status, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path("/v3/routeplanner/status")

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: route planner api: failed to get status: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: route planner api: failed to get status: unexpected status code %d received", resp.StatusCode)
	}

	var status routeplanner.Status
	if err := resp.JSON(&status); err != nil {
		return nil, fmt.Errorf("waterlink: rest: route planner api: failed to get status: %w", err)
	}

	return &status, nil
}

// FreeFailedAddress unmarks a failed address from the route planner. It
// dispatches a POST request to /v3/routeplanner/free/address.
func (cl *Client) FreeFailedAddress(address string) error {
	type reqBody struct {
		Address string `json:"address,omitempty"`
	}

	req := cl.client.Request().
		Method(http.MethodPost).
		Path("/v3/routeplanner/free/address")

	resp, err := req.JSON(&reqBody{Address: address}).Send()
	if err != nil {
		return fmt.Errorf("waterlink: rest: route planner api: failed to free address: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("waterlink: rest: route planner api: failed to free address: unexpected status code %d received", resp.StatusCode)
	}

	return nil
}

// FreeAllFailedAddresses unmarks all failed addresses from the route planner.
// It dispatches a POST request to /v3/routeplanner/free/all.
func (cl *Client) FreeAllFailedAddresses() error {
	req := cl.client.Request().
		Method(http.MethodPost).
		Path("/v3/routeplanner/free/address")

	resp, err := req.Send()
	if err != nil {
		return fmt.Errorf("waterlink: rest: route planner api: failed to free all: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("waterlink: rest: route planner api: failed to free all: unexpected status code %d received", resp.StatusCode)
	}

	return nil
}

// UpdateSession updates a session and returns the updated session. It
// dispatches a PATCH request to /v3/sessions/{sessionID}.
func (cl *Client) UpdateSession(sessionID string, update lavalink.Session) (*lavalink.Session, error) {
	req := cl.client.Request().
		Method(http.MethodPatch).
		Path(fmt.Sprintf("/v3/sessions/%s", sessionID)).
		Use(body.JSON(update))

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: session api: failed to update session: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: session api: failed to update session: unexpected status code %d received", resp.StatusCode)
	}

	var session lavalink.Session
	if err := resp.JSON(&session); err != nil {
		return nil, fmt.Errorf("waterlink: rest: session api: failed to update session: %w", err)
	}

	return &session, nil
}

// LoadTracks pre-loads one or multiple tracks using the given identifier. It
// dispatches a GET request to /v3/loadtracks?identifier={identifier}.
//
// See: player.Query
func (cl *Client) LoadTracks(identifier player.Query) (*player.TrackLoadResult, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path("/v3/loadtracks").
		SetQuery("identifier", string(identifier))

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to load track: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to load track: unexpected status code %d received", resp.StatusCode)
	}

	var result player.TrackLoadResult
	if err := resp.JSON(&result); err != nil {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to load track: %w", err)
	}

	return &result, nil
}

// DecodeTrack decodes the given player.EncodedTrack to a player.Track. It
// dispatches a GET request to /v3/decodetrack?encodedTrack={encodedTrack}.
func (cl *Client) DecodeTrack(encodedTrack player.EncodedTrack) (*player.Track, error) {
	req := cl.client.Request().
		Method(http.MethodGet).
		Path("/v3/decodetrack").
		SetQuery("encodedTrack", string(encodedTrack))

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to decode track: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to decode track: unexpected status code %d received", resp.StatusCode)
	}

	var track player.Track
	if err := resp.JSON(&track); err != nil {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to decode track: %w", err)
	}

	return &track, nil
}

// DecodeTracks decodes an array of player.EncodedTrack into an array of
// player.Track. It dispatches a POST request to /v3/decodetracks with the
// encoded tracks marshalled in the request body.
func (cl *Client) DecodeTracks(encodedTracks []player.EncodedTrack) ([]player.Track, error) {
	req := cl.client.Request().
		Method(http.MethodPost).
		Path("/v3/decodetracks").
		Use(body.JSON(encodedTracks))

	resp, err := req.Send()
	if err != nil {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to decode tracks: %w", err)
	}

	if !resp.Ok {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to decode tracks: unexpected status code %d received", resp.StatusCode)
	}

	var tracks []player.Track
	if err := resp.JSON(&tracks); err != nil {
		return nil, fmt.Errorf("waterlink: rest: track api: failed to decode tracks: %w", err)
	}

	return tracks, nil
}
