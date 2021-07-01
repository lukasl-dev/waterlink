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

	"github.com/lukasl-dev/waterlink/usecase/loadtrack"
)

const (
	loadTrackPath       = "/loadtracks"
	identifierParameter = "identifier"
)

type trackLoader struct {
	client     *http.Client
	host       url.URL
	passphrase string
}

func NewTrackLoader(client *http.Client, host url.URL, passphrase string) loadtrack.TrackLoader {
	host.Path += loadTrackPath
	return &trackLoader{
		client:     client,
		host:       host,
		passphrase: passphrase,
	}
}

func (l *trackLoader) LoadTrack(identifier string) (*loadtrack.Response, error) {
	resp, err := l.request(identifier)
	if err != nil {
		return nil, err
	}
	return l.unmarshal(resp)
}

func (l *trackLoader) unmarshal(resp *http.Response) (res *loadtrack.Response, err error) {
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(b, &res)
}

func (l *trackLoader) request(identifier string) (*http.Response, error) {
	req, err := l.createRequest(identifier)
	if err != nil {
		return nil, err
	}
	return l.client.Do(req)
}

func (l *trackLoader) createRequest(identifier string) (*http.Request, error) {
	host := l.createURL(identifier)
	req, err := http.NewRequest(http.MethodGet, host.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header = createHeader(l.passphrase)
	return req, nil
}

func (l *trackLoader) createURL(identifier string) url.URL {
	host := l.host
	host.RawQuery = l.createQuery(identifier).Encode()
	return host
}

func (l *trackLoader) createQuery(identifier string) url.Values {
	query := make(url.Values)
	query.Set(identifierParameter, identifier)
	return query
}
