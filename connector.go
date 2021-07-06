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
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/driver/websocketdriver"
)

type connector struct {
	opts *ConnectOptions
}

func Connect(ctx context.Context, host url.URL, opts ...*ConnectOptions) (Connection, error) {
	return newConnector(opts).open(ctx, host)
}

func newConnector(opts []*ConnectOptions) *connector {
	return &connector{
		opts: minimizeConnectOptions(opts),
	}
}

func (c *connector) open(ctx context.Context, host url.URL) (Connection, error) {
	conn, resp, err := c.dial(ctx, host)
	if err != nil {
		return nil, err
	}
	return c.connection(conn, resp), nil
}

func (c *connector) dial(ctx context.Context, host url.URL) (*websocket.Conn, *http.Response, error) {
	return c.opts.dialer.DialContext(ctx, host.String(), c.requestHeader())
}

func (c *connector) requestHeader() http.Header {
	h := make(http.Header)
	c.appendAuthHeader(&h)
	c.appendNumShardsHeader(&h)
	c.appendUserIDHeader(&h)
	c.appendResumeKeyHeader(h)
	return h
}

func (c *connector) appendAuthHeader(h *http.Header) {
	h.Set(headerAuthorization, c.opts.passphrase)
}

func (c *connector) appendNumShardsHeader(h *http.Header) {
	h.Set(headerNumShards, strconv.Itoa(int(c.opts.numShards)))
}

func (c *connector) appendUserIDHeader(h *http.Header) {
	h.Set(headerUserID, strconv.Itoa(int(c.opts.userID)))
}

func (c *connector) appendResumeKeyHeader(h http.Header) {
	if len(c.opts.resumeKey) > 0 {
		h.Set(headerResumeKey, c.opts.resumeKey)
	}
}

func (c *connector) connection(conn *websocket.Conn, resp *http.Response) Connection {
	return &connection{
		ResumingConfigurer: websocketdriver.NewResumeConfigurer(conn),
		Destroyer:          websocketdriver.NewDestroyer(conn),
		Equalizer:          websocketdriver.NewEqualizer(conn),
		Pauser:             websocketdriver.NewPauser(conn),
		Player:             websocketdriver.NewPlayer(conn),
		Seeker:             websocketdriver.NewSeeker(conn),
		Stopper:            websocketdriver.NewStopper(conn),
		VoiceUpdater:       websocketdriver.NewVoiceUpdater(conn),
		VolumeUpdater:      websocketdriver.NewVolumeUpdater(conn),
		resumed:            c.resumed(resp.Header),
	}
}

func (c *connector) resumed(h http.Header) bool {
	return h.Get(headerSessionResumed) == "true"
}
