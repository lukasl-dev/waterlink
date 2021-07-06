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
	"github.com/lukasl-dev/waterlink/usecase/updatevoice"
)

type voiceUpdater struct {
	conn *websocket.Conn
}

var _ updatevoice.VoiceUpdater = (*voiceUpdater)(nil)

func NewVoiceUpdater(conn *websocket.Conn) updatevoice.VoiceUpdater {
	return &voiceUpdater{
		conn: conn,
	}
}

type (
	voiceUpdatePayload struct {
		OP        op               `json:"op,omitempty"`
		GuildID   string           `json:"guildId,omitempty"`
		SessionID string           `json:"sessionId,omitempty"`
		Event     voiceUpdateEvent `json:"event,omitempty"`
	}
	voiceUpdateEvent struct {
		GuildID  string `json:"guild_id"`
		Token    string `json:"token"`
		Endpoint string `json:"endpoint"`
	}
)

func (u *voiceUpdater) UpdateVoice(guildID uint, sessionID, token, endpoint string) error {
	s := strconv.Itoa(int(guildID))
	return u.conn.WriteJSON(voiceUpdatePayload{
		OP:        opVoiceUpdate,
		GuildID:   s,
		SessionID: sessionID,
		Event: voiceUpdateEvent{
			GuildID:  s,
			Token:    token,
			Endpoint: endpoint,
		},
	})
}
