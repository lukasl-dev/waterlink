/*
 * MIT License
 *
 * Copyright (c) 2021 lukas.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package websocketdriver

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/adapter/receiveevent"
	"github.com/lukasl-dev/waterlink/entity/event"
	playerentity "github.com/lukasl-dev/waterlink/entity/player"
	"github.com/lukasl-dev/waterlink/entity/server"
)

type eventReceiver struct {
	conn   *websocket.Conn
	events chan event.Event
}

var _ receiveevent.EventReceiver = (*eventReceiver)(nil)

// NewEventReceiver returns a new event receiver.
func NewEventReceiver(conn *websocket.Conn) receiveevent.EventReceiver {
	return &eventReceiver{
		conn:   conn,
		events: make(chan event.Event),
	}
}

// Events returns a channel in which all events are streamed.
func (e *eventReceiver) Events() <-chan event.Event {
	go e.start()
	return e.events
}

func (e *eventReceiver) start() {
	for {
		evt, err := e.unmarshal()
		if err == nil {
			e.events <- evt
		}
	}
}

type genericEvent struct {
	Op      string `json:"op,omitempty"`
	Type    string `json:"type,omitempty"`
	GuildID string `json:"guildId,omitempty"`
	playerentity.TrackEnd
	playerentity.TrackException
	playerentity.TrackStart
	playerentity.TrackStuck
	playerentity.Update
	server.WebsocketClosed
	server.Stats
}

func (e *eventReceiver) unmarshal() (event.Event, error) {
	var payload genericEvent
	if err := e.conn.ReadJSON(&payload); err != nil {
		return nil, err
	}
	return e.lookupEvent(payload)
}

func (e *eventReceiver) lookupEvent(evt genericEvent) (event.Event, error) {
	if evt.Op == "stats" {
		return evt.Stats, nil
	}
	switch evt.Type {
	case "TrackEndEvent":
		evt.TrackEnd.GuildID = evt.GuildID
		return evt.TrackEnd, nil
	case "TrackExceptionEvent":
		evt.TrackException.GuildID = evt.GuildID
		return evt.TrackException, nil
	case "TrackStartEvent":
		evt.TrackStart.GuildID = evt.GuildID
		return evt.TrackStart, nil
	case "TrackStuckEvent":
		evt.TrackStuck.GuildID = evt.GuildID
		return evt.TrackStuck, nil
	case "WebSocketClosedEvent":
		evt.WebsocketClosed.GuildID = evt.GuildID
		return evt.WebsocketClosed, nil
	default:
		return nil, errors.New("unknown event received")
	}
}
