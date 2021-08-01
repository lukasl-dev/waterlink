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

package waterlink

import (
	"github.com/lukasl-dev/waterlink/entity/event"
	"github.com/lukasl-dev/waterlink/usecase/equalize"
	"github.com/lukasl-dev/waterlink/usecase/play"
	"github.com/stretchr/testify/mock"
)

// MockedConnection is the mocking implementation of
// Connection.
type MockedConnection struct {
	mock.Mock
}

var _ Connection = (*MockedConnection)(nil)

// Events returns a channel in which all events are streamed.
func (conn *MockedConnection) Events() <-chan event.Event {
	return conn.Called().Get(0).(<-chan event.Event)
}

// ConfigureResuming configures the resume key used
// to resume a connection.
func (conn *MockedConnection) ConfigureResuming(resumeKey string, timeout uint) error {
	return conn.Called(resumeKey, timeout).Error(0)
}

// Destroy is used to destroy a guild's audio player.
func (conn *MockedConnection) Destroy(guildID string) error {
	return conn.Called(guildID).Error(0)
}

// UseEqualizer applies the passed bands on a guild's
// audio player.
func (conn *MockedConnection) UseEqualizer(guildID string, bands ...equalize.Band) error {
	return conn.Called(guildID, bands).Error(0)
}

// SetPaused sets the paused state of a guild's
// audio player to the passed parameter value.
func (conn *MockedConnection) SetPaused(guildID string, paused bool) error {
	return conn.Called(guildID, paused).Error(0)
}

// Play plays the track with the given id on the
// guild's audio player.
// More options can be configured via Options.
func (conn *MockedConnection) Play(guildID string, trackID string, opts ...*play.Options) error {
	return conn.Called(guildID, trackID, opts).Error(0)
}

// Seek skips the current track of a guild's audio
// player to the passed position.
func (conn *MockedConnection) Seek(guildID string, position uint) error {
	return conn.Called(guildID, position).Error(0)
}

// Stop stops the current track of a guild's audio
// player.
func (conn *MockedConnection) Stop(guildID string) error {
	return conn.Called(guildID).Error(0)
}

// UpdateVoice is sent when the voice server of a guild
// has been updated.
// This method must be performed to play a track.
// See: https://discord.com/developers/docs/topics/gateway#voice-server-update
func (conn *MockedConnection) UpdateVoice(guildID string, sessionID, token, endpoint string) error {
	return conn.Called(guildID, sessionID, token, endpoint).Error(0)
}

// UpdateVolume changes the volume of a guild's
// audio player.
func (conn *MockedConnection) UpdateVolume(guildID string, volume uint) error {
	return conn.Called(guildID, volume).Error(0)
}

// Resumed returns true whenever the Connection has
// been resumed.
func (conn *MockedConnection) Resumed() bool {
	return conn.Called().Bool(0)
}
