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
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNewConnectOptions(t *testing.T) {
	actual := NewConnectOptions()
	expected := &ConnectOptions{
		dialer: websocket.DefaultDialer,
	}
	assert.Equal(t, expected, actual)
}

func TestConnectOptions_WithDialer(t *testing.T) {
	var (
		dialer = &websocket.Dialer{}
	)
	actual := NewConnectOptions().WithDialer(dialer)
	expected := &ConnectOptions{
		dialer: dialer,
	}
	assert.Equal(t, expected, actual)
}

func TestConnectOptions_WithPassphrase(t *testing.T) {
	const (
		passphrase = "youshallnotpass"
	)
	actual := NewConnectOptions().WithPassphrase(passphrase)
	expected := &ConnectOptions{
		dialer:     websocket.DefaultDialer,
		passphrase: passphrase,
	}
	assert.Equal(t, expected, actual)
}

func TestConnectOptions_WithNumShards(t *testing.T) {
	const (
		numShards = 12
	)
	actual := NewConnectOptions().WithNumShards(numShards)
	expected := &ConnectOptions{
		dialer:    websocket.DefaultDialer,
		numShards: numShards,
	}
	assert.Equal(t, expected, actual)
}

func TestConnectOptions_WithUserID(t *testing.T) {
	const (
		userID = 432617716961116180
	)
	actual := NewConnectOptions().WithUserID(userID)
	expected := &ConnectOptions{
		dialer: websocket.DefaultDialer,
		userID: userID,
	}
	assert.Equal(t, expected, actual)
}

func TestConnectOptions_WithResumeKey(t *testing.T) {
	const (
		resumeKey = "myResumeKey"
	)
	actual := NewConnectOptions().WithResumeKey(resumeKey)
	expected := &ConnectOptions{
		dialer:    websocket.DefaultDialer,
		resumeKey: resumeKey,
	}
	assert.Equal(t, expected, actual)
}
