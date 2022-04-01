package waterlink

import (
	"encoding/json"
	"fmt"
	"github.com/gompus/snowflake"
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/event"
	"github.com/lukasl-dev/waterlink/internal/message"
	"github.com/lukasl-dev/waterlink/internal/message/opcode"
	"github.com/lukasl-dev/waterlink/internal/pkgerror"
	"net/http"
	"time"
)

// Connection is used to receive and dispatch messages from the
type Connection struct {
	// opts contains optional values for this connection.
	opts ConnectionOptions

	// conn is the underlying websocket connection.
	conn *websocket.Conn

	// sessionResumed reports whether a previous session has been resumed.
	sessionResumed bool

	// apiVersion is server's API version that was acquired during the handshake.
	apiVersion string
}

// Open opens a new websocket connection to addr. The given creds are used to
// authenticate the connection to use the protocol.
func Open(addr string, creds Credentials, opts ...ConnectionOptions) (*Connection, error) {
	switch {
	case len(opts) == 0:
		opts = []ConnectionOptions{defaultConnectionOptions}
	case len(opts) > 1:
		panic(pkgerror.New("connection: open: too many options"))
	}

	conn, resp, err := websocket.DefaultDialer.Dial(addr, creds.header())
	if err != nil {
		return nil, pkgerror.Wrap("connection: open", err)
	}

	return wrapConn(opts[0], conn, resp), nil
}

// wrapConn wraps the given websocket connection.
func wrapConn(opts ConnectionOptions, conn *websocket.Conn, resp *http.Response) *Connection {
	c := &Connection{opts: opts, conn: conn}
	if opts.EventBus != nil {
		go c.listen()
	}
	c.sessionResumed = resp.Header.Get("Session-Resumed") == "true"
	c.apiVersion = resp.Header.Get("Lavalink-Api-Version")
	return c
}

// SessionResumed returns true whether a previous session has been resumed.
func (conn Connection) SessionResumed() bool {
	return conn.sessionResumed
}

// APIVersion returns the server's API version that was acquired during the
// handshake.
func (conn Connection) APIVersion() string {
	return conn.apiVersion
}

// Guild returns a Guild used to interact with a specific guild. The
// availability is not checked client-side.
func (conn Connection) Guild(id snowflake.Snowflake) Guild {
	return Guild{w: conn.conn, id: id}
}

// ConfigureResuming enable the resumption of the session and defines the number
// of seconds after which the session will be considered expired server-side. This
// is useful to avoid stopping the audio players that are related to the session.
func (conn Connection) ConfigureResuming(key string, timeout time.Duration) error {
	return pkgerror.Wrap("connection: configure resuming", conn.conn.WriteJSON(
		message.ConfigureResuming{
			Outgoing: message.Outgoing{Op: opcode.ConfigureResuming},
			Key:      &key,
			Timeout:  uint(timeout.Seconds()),
		},
	))
}

// DisableResuming disables the resumption of the session. If disabled, audio
// players will stop immediately after the connection is closed.
func (conn Connection) DisableResuming() error {
	return pkgerror.Wrap("connection: disable resuming", conn.conn.WriteJSON(
		message.ConfigureResuming{
			Outgoing: message.Outgoing{Op: opcode.ConfigureResuming},
		},
	))
}

func (conn Connection) listen() {
	for {
		op, data, err := conn.readMessage()
		if err != nil {
			continue
		}
		_ = conn.emit(op, data)
	}
}

func (conn Connection) readMessage() (op string, data []byte, err error) {
	_, data, err = conn.conn.ReadMessage()
	if err != nil {
		return "", nil, pkgerror.Wrap("connection: event bus: listen: read message", err)
	}

	op, err = conn.opcodeOf(data)
	if err != nil {
		return "", nil, pkgerror.Wrap("connection: event bus: listen: invalid message received", err)
	}

	return op, data, nil
}

func (conn Connection) opcodeOf(data []byte) (string, error) {
	var msg message.Incoming
	if err := json.Unmarshal(data, &msg); err != nil {
		return "", err
	}

	switch msg.Op {
	case opcode.PlayerUpdate, opcode.Stats, opcode.Event:
		return string(msg.Op), nil
	default:
		return "", fmt.Errorf("unknown opcode %q", msg.Op)
	}
}

func (conn Connection) emit(op string, data []byte) error {
	switch opcode.Incoming(op) {
	case opcode.PlayerUpdate:
		return conn.emitPlayerUpdate(data)
	case opcode.Stats:
		return conn.emitStats(data)
	case opcode.Event:
		return conn.emitEvent(data)
	default:
		return fmt.Errorf("unknown opcode %q", op)
	}
}

func (conn Connection) emitPlayerUpdate(data []byte) error {
	var e event.PlayerUpdate
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	conn.opts.EventBus.EmitPlayerUpdate(e)
	return nil
}

func (conn Connection) emitStats(data []byte) error {
	var e event.Stats
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	conn.opts.EventBus.EmitStats(e)
	return nil
}

func (conn Connection) emitEvent(data []byte) error {
	var msg message.Event
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}

	switch msg.Type {
	case message.EventTypeTrackStart:
		return conn.emitTrackStart(data)
	case message.EventTypeTrackEnd:
		return conn.emitTrackEnd(data)
	case message.EventTypeTrackException:
		return conn.emitTrackException(data)
	case message.EventTypeTrackStuck:
		return conn.emitTrackStuck(data)
	case message.EventTypeWebSocketClosed:
		return conn.emitWebSocketClosed(data)
	default:
		return fmt.Errorf("unknown event type %q", msg.Type)
	}
}

func (conn Connection) emitTrackStart(data []byte) error {
	var e event.TrackStart
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	conn.opts.EventBus.EmitTrackStart(e)
	return nil
}

func (conn Connection) emitTrackEnd(data []byte) error {
	var e event.TrackEnd
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	conn.opts.EventBus.EmitTrackEnd(e)
	return nil
}

func (conn Connection) emitTrackException(data []byte) error {
	var e event.TrackException
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	conn.opts.EventBus.EmitTrackException(e)
	return nil
}

func (conn Connection) emitTrackStuck(data []byte) error {
	var e event.TrackStuck
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	conn.opts.EventBus.EmitTrackStuck(e)
	return nil
}

func (conn Connection) emitWebSocketClosed(data []byte) error {
	var e event.WebSocketClosed
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	conn.opts.EventBus.EmitWebSocketClosed(e)
	return nil
}
