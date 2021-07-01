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
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/usecase/seek"
)

type seekPayload struct {
	OP       op     `json:"op,omitempty"`
	GuildID  string `json:"guildId,omitempty"`
	Position uint   `json:"position,omitempty"`
}

type seeker struct {
	conn *websocket.Conn
}

func NewSeeker(conn *websocket.Conn) seek.Seeker {
	return &seeker{
		conn: conn,
	}
}

func (s *seeker) Seek(guildID string, position uint) error {
	return s.conn.WriteJSON(seekPayload{
		OP:       opSeek,
		GuildID:  guildID,
		Position: position,
	})
}