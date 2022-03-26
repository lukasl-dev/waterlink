package waterlink

import (
	"github.com/gompus/snowflake"
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/internal/pkgerror"
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
