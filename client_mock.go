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
	"github.com/lukasl-dev/waterlink/usecase/loadtrack"
	"github.com/lukasl-dev/waterlink/usecase/routeplanner/getstatus"
	"github.com/stretchr/testify/mock"
)

type MockedClient struct {
	mock.Mock
}

func NewMockedClient() *MockedClient {
	return new(MockedClient)
}

func (c *MockedClient) DecodeTracks(trackIDs ...string) ([]*loadtrack.Track, error) {
	args := c.Called(trackIDs)
	return args.Get(0).([]*loadtrack.Track), args.Error(1)
}

func (c *MockedClient) LoadTrack(identifier string) (*loadtrack.Response, error) {
	args := c.Called(identifier)
	return args.Get(0).(*loadtrack.Response), args.Error(1)
}

func (c *MockedClient) Status() (*getstatus.Status, error) {
	args := c.Called()
	return args.Get(0).(*getstatus.Status), args.Error(1)
}

func (c *MockedClient) UnmarkAddress(address string) error {
	return c.Called(address).Error(0)
}

func (c *MockedClient) UnmarkAddresses() error {
	return c.Called().Error(0)
}
