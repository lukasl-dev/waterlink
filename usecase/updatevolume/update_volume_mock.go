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

package updatevolume

import "github.com/stretchr/testify/mock"

// MockedVolumeUpdater is the mock implementation of VolumeUpdater.
type MockedVolumeUpdater struct {
	mock.Mock
}

var _ VolumeUpdater = (*MockedVolumeUpdater)(nil)

// NewMockedVolumeUpdater returns a new MockedVolumeUpdater.
func NewMockedVolumeUpdater() *MockedVolumeUpdater {
	return new(MockedVolumeUpdater)
}

// UpdateVolume changes the volume of a guild's
// audio player.
func (u *MockedVolumeUpdater) UpdateVolume(guildID, volume uint) error {
	return u.Called(guildID, volume).Error(0)
}
