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

import "github.com/gorilla/websocket"

// ConnectOptions is used to configure further specifications
// of the Connect method.
type ConnectOptions struct {
	dialer     *websocket.Dialer
	passphrase string
	numShards  uint
	userID     uint
	resumeKey  string
}

// NewConnectOptions returns a new ConnectOptions.
func NewConnectOptions() *ConnectOptions {
	return &ConnectOptions{
		dialer: websocket.DefaultDialer,
	}
}

// minimizeConnectOptions minimizes the passed options
// to a single one.
func minimizeConnectOptions(opts []*ConnectOptions) *ConnectOptions {
	if len(opts) > 0 {
		return opts[0]
	}
	return NewConnectOptions()
}

// WithDialer sets the dialer to the parameter value.
func (opts *ConnectOptions) WithDialer(dialer *websocket.Dialer) *ConnectOptions {
	opts.dialer = dialer
	return opts
}

// WithPassphrase sets the passphrase to the parameter value.
func (opts *ConnectOptions) WithPassphrase(passphrase string) *ConnectOptions {
	opts.passphrase = passphrase
	return opts
}

// WithNumShards sets the number of shards to the parameter value.
func (opts *ConnectOptions) WithNumShards(numShards uint) *ConnectOptions {
	opts.numShards = numShards
	return opts
}

// WithUserID sets the user id to the parameter value.
func (opts *ConnectOptions) WithUserID(userID uint) *ConnectOptions {
	opts.userID = userID
	return opts
}

// WithResumeKey sets the resume key to the parameter value.
func (opts *ConnectOptions) WithResumeKey(resumeKey string) *ConnectOptions {
	opts.resumeKey = resumeKey
	return opts
}
