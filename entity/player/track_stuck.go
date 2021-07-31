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

// TrackStuck is fired when a track was started, but no audio frames from it
// have arrived in a long time, specified by the server-side defined threshold.
type TrackStuck struct {
	GuildID     uint   `json:"guildId,omitempty"`
	TrackID     string `json:"track,omitempty"`
	ThresholdMS uint   `json:"thresholdMs,omitempty"`
}

var _ event.Event = (*TrackStuck)(nil)

// Type returns the TrackStuck event's type that corresponds to the
// constant event.TrackStuck.
func (t TrackStuck) Type() event.Type {
	return event.TrackStuck
}
