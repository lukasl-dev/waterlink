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
	"github.com/lukasl-dev/waterlink/entity/routeplanner"
	"github.com/lukasl-dev/waterlink/entity/track"
	"github.com/lukasl-dev/waterlink/usecase/loadtracks"
	"github.com/stretchr/testify/mock"
)

// MockedRequester is the mocking implementation of Requester.
type MockedRequester struct {
	mock.Mock
}

var _ Requester = (*MockedRequester)(nil)

// NewMockedRequester returns a new MockedRequester.
func NewMockedRequester() *MockedRequester {
	return new(MockedRequester)
}

// DecodeTracks is used to decode the passed trackIDs
// to track infos.
func (r *MockedRequester) DecodeTracks(trackIDs ...string) ([]*track.Info, error) {
	args := r.Called(trackIDs)
	return args.Get(0).([]*track.Info), args.Error(1)
}

// LoadTracks loads multiple tracks by the passed
// identifier.
func (r *MockedRequester) LoadTracks(identifier string) (*loadtracks.Response, error) {
	args := r.Called(identifier)
	return args.Get(0).(*loadtracks.Response), args.Error(1)
}

// Status returns the routeplanner's status.
func (r *MockedRequester) Status() (*routeplanner.Status, error) {
	args := r.Called()
	return args.Get(0).(*routeplanner.Status), args.Error(1)
}

// UnmarkAddress unmarks the passed (failed) address.
func (r *MockedRequester) UnmarkAddress(address string) error {
	return r.Called(address).Error(0)
}

// UnmarkAddresses unmarks all failed addresses.
func (r *MockedRequester) UnmarkAddresses() error {
	return r.Called().Error(0)
}
