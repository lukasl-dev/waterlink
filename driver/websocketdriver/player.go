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
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/usecase/play"
)

type player struct {
	conn *websocket.Conn
}

var _ play.Player = (*player)(nil)

func NewPlayer(conn *websocket.Conn) play.Player {
	return &player{
		conn: conn,
	}
}

type playPayload struct {
	OP        op     `json:"op,omitempty"`
	GuildID   string `json:"guildId,omitempty"`
	Track     string `json:"track,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Volume    string `json:"volume,omitempty"`
	NoReplace bool   `json:"noReplace,omitempty"`
	Pause     bool   `json:"pause,omitempty"`
}

func (p *player) Play(guildID uint, trackID string, opts ...*play.Options) error {
	return p.conn.WriteJSON(p.payload(guildID, trackID, opts))
}

func (p *player) payload(guildID uint, trackID string, opts []*play.Options) playPayload {
	payload := playPayload{
		OP:      opPlay,
		GuildID: strconv.Itoa(int(guildID)),
		Track:   trackID,
	}
	p.insert(&payload, play.MinimizeOptions(opts...))
	return payload
}

func (p *player) insert(payload *playPayload, opts *play.Options) {
	payload.StartTime = strconv.Itoa(int(opts.StartTime))
	payload.EndTime = strconv.Itoa(int(opts.EndTime))
	payload.Volume = strconv.Itoa(opts.Volume)
	payload.NoReplace = opts.NoReplace
	payload.Pause = opts.Paused
}
