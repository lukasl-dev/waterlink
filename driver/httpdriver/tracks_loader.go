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

package httpdriver

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/lukasl-dev/waterlink/usecase/loadtracks"
)

const (
	pathLoadTracks  = "/loadtracks"
	paramIdentifier = "identifier"
)

type tracksLoader struct {
	client     *http.Client
	host       url.URL
	passphrase string
}

var _ loadtracks.TracksLoader = (*tracksLoader)(nil)

// NewTrackLoader returns a new TrackLoader.
func NewTrackLoader(client *http.Client, host url.URL, passphrase string) loadtracks.TracksLoader {
	host.Path += pathLoadTracks
	return &tracksLoader{
		client: client,
		host:   host, passphrase: passphrase,
	}
}

// LoadTracks loads multiple tracks by the passed
// identifier.
func (l *tracksLoader) LoadTracks(identifier string) (*loadtracks.Response, error) {
	req, err := l.request(identifier)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(req)
	if err != nil {
		return nil, err
	}
	return l.unmarshal(resp)
}

func (l *tracksLoader) request(identifier string) (*http.Request, error) {
	host := l.host
	host.RawQuery = l.query(identifier).Encode()
	req, err := http.NewRequest(http.MethodGet, host.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header = authHeader(l.passphrase)
	return req, nil
}

func (l *tracksLoader) query(identifier string) url.Values {
	q := make(url.Values)
	q.Set(paramIdentifier, identifier)
	return q
}

func (l *tracksLoader) unmarshal(resp *http.Response) (dest *loadtracks.Response, err error) {
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return dest, json.Unmarshal(b, &dest)
}
