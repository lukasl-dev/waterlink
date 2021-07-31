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

package loadtrack

import "github.com/stretchr/testify/mock"

// MockedTrackLoader is the mocking implementation of
// TrackLoader.
type MockedTrackLoader struct {
	mock.Mock
}

var _ TrackLoader = (*MockedTrackLoader)(nil)

// NewMockedTrackLoader returns a new MockedTrackLoader.
func NewMockedTrackLoader() *MockedTrackLoader {
	return new(MockedTrackLoader)
}

// LoadTracks loads multiple tracks by the passed
// identifier.
func (l *MockedTrackLoader) LoadTracks(identifier string) (*Response, error) {
	args := l.Called(identifier)
	return args.Get(0).(*Response), args.Error(1)
}
