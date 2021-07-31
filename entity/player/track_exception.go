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

package player

import "github.com/lukasl-dev/waterlink/entity/event"

// TrackException is fired when an exception occurs in an audio track that
// causes it to halt or not start.
type TrackException struct {
	// GuildID of the guild on which the audio player is operating on.
	GuildID uint `json:"guildId,omitempty"`

	// TrackID of the track that halt or not start.
	TrackID string `json:"track,omitempty"`

	// Error is the error message.
	Error string `json:"error,omitempty"`
}

var _ event.Event = (*TrackException)(nil)

// Type returns the TrackStuck event's type that corresponds to the
// constant event.TrackException.
func (t TrackException) Type() event.Type {
	return event.TrackException
}
