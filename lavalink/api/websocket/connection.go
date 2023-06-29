package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/v3/internal"
)

// Connection is a wrapper for an underlying websocket connection for a more
// convenient usage of the Lavalink websocket API.
type Connection struct {
	// ws is the underlying websocket connection. If nil, the connection is
	// closed.
	ws *websocket.Conn

	// ro is the Router which is responsible for handling incoming events.
	ro *Router

	// listening indicates whether the websocket connection is currently
	// listening for incoming messages.
	listening bool
}

var (
	// ErrConnectionInvalidAddress occurs when an invalid address is passed to
	// Open().
	//
	// Valid address:
	// ws://localhost:2333
	//
	// Invalid address:
	// localhost:2333
	ErrConnectionInvalidAddress = fmt.Errorf("waterlink: rest: address is invalid")

	// ErrConnectionClosed occurs when an operation, which requires an active
	// connection, is performed on a closed connection.
	ErrConnectionClosed = fmt.Errorf("waterlink: websocket: connection already closed")

	// ErrConnectionListening occurs when an operation, which requires a
	// non-listening connection, is performed on a listening connection.
	ErrConnectionListening = fmt.Errorf("waterlink: websocket: connection already listening")
)

// Open opens a new websocket connection to the given address with the given
// header. It returns an error if the connection could not be established.
//
// Example address: ws://localhost:2333
func Open(addr string, opt Options) (*Connection, error) {
	if !internal.IsURL(addr) {
		return nil, ErrConnectionInvalidAddress
	}

	ws, _, err := websocket.DefaultDialer.Dial(addr, opt.http())
	if err != nil {
		return nil, fmt.Errorf("waterlink: websocket: %w", err)
	}

	return &Connection{ws: ws, ro: newMessageRouter()}, nil
}

// Open returns whether the websocket connection is open or not.
func (conn *Connection) Open() bool {
	return conn.ws != nil
}

// Close closes the websocket connection. It returns an error if the connection
// could not be closed or if the connection is already closed.
func (conn *Connection) Close() error {
	if conn.ws == nil {
		return ErrConnectionClosed
	}

	err := conn.ws.Close()
	if err != nil {
		return fmt.Errorf("waterlink: websocket: %w", err)
	}

	conn.ws = nil
	return nil
}

// Router returns the Router of the websocket connection. It returns
// an error if the connection is already closed.
func (conn *Connection) Router() (*Router, error) {
	if conn.ws == nil {
		return nil, ErrConnectionClosed
	}

	return conn.ro, nil
}

// Listen starts listening for incoming messages from the websocket connection.
// It returns a channel that receives errors as soon as they occur.
func (conn *Connection) Listen() (<-chan error, error) {
	if conn.ws == nil {
		return nil, ErrConnectionClosed
	}

	if conn.listening {
		return nil, ErrConnectionListening
	}
	conn.listening = true

	errs := make(chan error)
	go func() {
		for conn.ws != nil {
			_, data, err := conn.ws.ReadMessage()
			if err != nil {
				errs <- fmt.Errorf("waterlink: websocket: %w", err)
			}

			err = conn.ro.route(data)
			if err != nil {
				errs <- fmt.Errorf("waterlink: websocket: %w", err)
			}
		}
	}()

	return errs, nil
}
