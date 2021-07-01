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
	"github.com/lukasl-dev/waterlink/usecase/play"
)

type playPayload struct {
	OP        op     `json:"op,omitempty"`
	GuildID   string `json:"guildId,omitempty"`
	Track     string `json:"track,omitempty"`
	StartTime uint   `json:"startTime,omitempty"`
	EndTime   uint   `json:"endTime,omitempty"`
	Volume    int    `json:"volume,omitempty"`
	NoReplace bool   `json:"noReplace,omitempty"`
	Pause     bool   `json:"pause,omitempty"`
}

type player struct {
	conn *websocket.Conn
}

func NewPlayer(conn *websocket.Conn) play.Player {
	return &player{
		conn: conn,
	}
}

func (p *player) Play(guildID string, trackID string, opts *play.Options) error {
	return p.conn.WriteJSON(p.createPayload(guildID, trackID, opts))
}

func (p *player) createPayload(guildID string, trackID string, opts *play.Options) playPayload {
	payload := playPayload{
		OP:      opPlay,
		GuildID: guildID,
		Track:   trackID,
	}
	p.insertOptions(&payload, opts)
	return payload
}

func (p *player) insertOptions(payload *playPayload, opts *play.Options) {
	if opts != nil {
		payload.StartTime = opts.StartTime
		payload.EndTime = opts.EndTime
		payload.Volume = opts.Volume
		payload.NoReplace = opts.NoReplace
		payload.Pause = opts.Pause
	}
}
