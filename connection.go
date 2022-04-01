package waterlink

import (
	"github.com/gompus/snowflake"
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/internal/message"
	"github.com/lukasl-dev/waterlink/internal/message/opcode"
	"github.com/lukasl-dev/waterlink/internal/pkgerror"
	"time"
)

// Connection is used to receive and dispatch messages from the
type Connection struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// Open opens a new websocket connection to addr. The given creds are used to
// authenticate the connection to use the protocol.
func Open(addr string, creds Credentials) (*Connection, error) {
	conn, _, err := websocket.DefaultDialer.Dial(addr, creds.header())
	if err != nil {
		return nil, pkgerror.Wrap("open", err)
	}
	return newConnection(conn), nil
}

// newConnection wraps the given websocket connection.
func newConnection(conn *websocket.Conn) *Connection {
	return &Connection{conn: conn}
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
