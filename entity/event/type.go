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

package event

// Type is used to differentiate Event types.
type Type string

const (
	// TrackEnd is fired when an audio track ends in an audio player,
	// either by interruption, exception or reaching the end.
	// See also: player.TrackEnd
	TrackEnd Type = "track-end"

	// TrackException is fired when an exception occurs in an audio track
	// that causes it to halt or not start.
	// See also: player.TrackException
	TrackException Type = "track-exception"

	// TrackStart is fired when a track starts playing.
	// See also: player.TrackStart
	TrackStart Type = "track-start"

	// TrackStuck is fired when a track was started, but no audio frames from
	// it have arrived in a long time, specified by the predefined threshold.
	// See also: player.TrackStuck
	TrackStuck Type = "track-stuck"

	// PlayerUpdate TODO
	// See also: player.Update
	PlayerUpdate Type = "player-update"

	// WebsocketClosed when an audio web socket (to Discord) is closed.
	// This can happen for various reasons (normal and abnormal), e.g when using
	// an expired voice server update.
	// See also: server.WebsocketClosed
	WebsocketClosed Type = "websocket-closed"

	// Stats is fired repetitively every minute during an active connection.
	// See also: server.Stats
	Stats Type = "stats"
)
