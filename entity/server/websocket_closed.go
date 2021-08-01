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

package server

import "github.com/lukasl-dev/waterlink/entity/event"

// WebsocketClosed is fired when an audio web socket (to Discord) is closed.
// This can happen for various reasons (normal and abnormal), e.g when using an
// expired voice server update.
type WebsocketClosed struct {
	// GuildID is the id of the guild on which this event occurred.
	GuildID string `json:"guildId,omitempty"`

	// Code is the code returned by the discord api.
	Code uint `json:"code,omitempty"`

	// Reason is the reason why the audio connection closed.
	Reason string `json:"reason,omitempty"`

	// Remote TODO
	Remote bool `json:"byRemote,omitempty"`
}

var _ event.Event = (*WebsocketClosed)(nil)

// Type returns the WebsocketClosed event's type that corresponds to the
// constant event.WebsocketClosed.
func (w WebsocketClosed) Type() event.Type {
	return event.WebsocketClosed
}
