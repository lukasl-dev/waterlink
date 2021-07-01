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
	"github.com/lukasl-dev/waterlink/usecase/equalize"
	"github.com/lukasl-dev/waterlink/usecase/play"
	"github.com/stretchr/testify/mock"
)

type MockedConnection struct {
	mock.Mock
}

func NewMockedConnection() *MockedConnection {
	return new(MockedConnection)
}

func (c *MockedConnection) Destroy(guildID string) error {
	return c.Called(guildID).Error(0)
}

func (c *MockedConnection) UseEqualizer(guildID string, bands ...equalize.Band) error {
	return c.Called(guildID, bands).Error(0)
}

func (c *MockedConnection) SetPaused(guildID string, paused bool) error {
	return c.Called(guildID, paused).Error(0)
}

func (c *MockedConnection) Play(guildID, trackID string, opts *play.Options) error {
	return c.Called(guildID, trackID, opts).Error(0)
}

func (c *MockedConnection) ResumeSession(resumeKey string, timeout uint) error {
	return c.Called(resumeKey, timeout).Error(0)
}

func (c *MockedConnection) Seek(guildID string, position uint) error {
	return c.Called(guildID, position).Error(0)
}

func (c *MockedConnection) Stop(guildID string) error {
	return c.Called(guildID).Error(0)
}

func (c *MockedConnection) UpdateVoice(guildID, sessionID, token, endpoint string) error {
	return c.Called(guildID, sessionID, token, endpoint).Error(0)
}

func (c *MockedConnection) UpdateVolume(guildID string, volume uint) error {
	return c.Called(guildID, volume).Error(0)
}
