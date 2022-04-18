package waterlink

import (
	"errors"
	"fmt"
	"github.com/gompus/snowflake"
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/internal/message"
	"github.com/lukasl-dev/waterlink/internal/message/opcode"
	"github.com/lukasl-dev/waterlink/internal/pkgerror"
	"net/http"
	"time"
)

// Connection is used to receive and dispatch messages from the
type Connection struct {
	opts ConnectionOptions

	// conn is the underlying websocket connection.
	conn *websocket.Conn

	// closed indicates whether the connection has been closed.
	closed bool

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

	return wrapConn(opts[0], conn, resp.Header), nil
}

// wrapConn wraps the given websocket connection.
func wrapConn(opts ConnectionOptions, conn *websocket.Conn, h http.Header) *Connection {
	c := &Connection{opts: opts, conn: conn}
	c.header(h)
	if opts.EventHandler != nil {
		go c.listenForEvents()
	}
	return c
}

// header adapts the given header's values to the connection's internal state.
func (conn *Connection) header(h http.Header) {
	conn.sessionResumed = h.Get("Session-Resumed") == "true"
	conn.apiVersion = h.Get("Lavalink-Api-Version")
}

// Closed returns whether the connection has been closed.
func (conn *Connection) Closed() bool {
	return conn.closed
}

// Close closes the underlying websocket connection.
func (conn *Connection) Close() error {
	conn.closed = true
	return conn.conn.Close()
}

// SessionResumed returns true whether a previous session has been resumed.
func (conn *Connection) SessionResumed() bool {
	return conn.sessionResumed
}

// APIVersion returns the server's API version that was acquired during the
// handshake.
func (conn *Connection) APIVersion() string {
	return conn.apiVersion
}

// Guild returns a Guild used to interact with a specific guild. The
// availability is not checked client-side.
func (conn *Connection) Guild(id snowflake.Snowflake) Guild {
	return Guild{conn: conn, id: id}
}

// ConfigureResuming enable the resumption of the session and defines the number
// of seconds after which the session will be considered expired server-side. This
// is useful to avoid stopping the audio players that are related to the session.
func (conn *Connection) ConfigureResuming(key string, timeout time.Duration) error {
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
func (conn *Connection) DisableResuming() error {
	return pkgerror.Wrap("connection: disable resuming", conn.conn.WriteJSON(
		message.ConfigureResuming{
			Outgoing: message.Outgoing{Op: opcode.ConfigureResuming},
		},
	))
}

// listenForEvents runs the event loop that reads incoming messages from the
// server and dispatches them to the given eventBus.
func (conn *Connection) listenForEvents() {
	bus := newEventBus(conn.opts.EventHandler)
	for {
		_, data, err := conn.conn.ReadMessage()
		if err != nil {
			conn.handleEventError(err)
			break
		}
		if err := bus.receive(data); err != nil {
			conn.handleEventError(err)
		}
	}
}

// handleEventError handles errors that occur during the event loop.
func (conn *Connection) handleEventError(err interface{}) {
	_ = conn.Close()
	if conn.opts.HandleEventError != nil {
		var issue error
		switch err := err.(type) {
		case error:
			issue = err
		default:
			issue = errors.New(fmt.Sprint(err))
		}
		conn.opts.HandleEventError(pkgerror.Wrap("connection", issue))
	}
}
