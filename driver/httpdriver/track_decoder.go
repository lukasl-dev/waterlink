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
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/lukasl-dev/waterlink/entity/track"
	"github.com/lukasl-dev/waterlink/usecase/decodetrack"
)

const decodeTracksPath = "/decodetracks"

type trackDecoder struct {
	client     *http.Client
	host       url.URL
	passphrase string
}

func NewTrackDecoder(client *http.Client, host url.URL, passphrase string) decodetrack.TrackDecoder {
	host.Path += decodeTracksPath
	return &trackDecoder{
		client:     client,
		host:       host,
		passphrase: passphrase,
	}
}

func (d *trackDecoder) DecodeTracks(trackIDs ...string) ([]*track.Info, error) {
	resp, err := d.request(trackIDs)
	if err != nil {
		return nil, err
	}
	return d.unmarshal(resp)
}

func (d *trackDecoder) unmarshal(resp *http.Response) (tracks []*track.Info, err error) {
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return tracks, json.Unmarshal(b, &tracks)
}

func (d *trackDecoder) request(trackIDs []string) (*http.Response, error) {
	req, err := d.createRequest(trackIDs)
	if err != nil {
		return nil, err
	}
	return d.client.Do(req)
}

func (d *trackDecoder) createRequest(trackIDs []string) (*http.Request, error) {
	body, err := d.createBody(trackIDs)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(http.MethodGet, d.host.String(), body)
}

func (d *trackDecoder) createBody(trackIDs []string) (io.Reader, error) {
	b, err := json.Marshal(trackIDs)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
