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

const pathDecodeTracks = "/decodetracks"

type trackDecoder struct {
	client     *http.Client
	host       url.URL
	passphrase string
}

var _ decodetrack.TrackDecoder = (*trackDecoder)(nil)

func NewTrackDecoder(client *http.Client, host url.URL, passphrase string) decodetrack.TrackDecoder {
	host.Path += pathDecodeTracks
	return &trackDecoder{
		client:     client,
		host:       host,
		passphrase: passphrase,
	}
}

func (d *trackDecoder) DecodeTracks(trackIDs ...string) ([]*track.Info, error) {
	req, err := d.request(trackIDs)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	return d.unmarshal(resp)
}

func (d *trackDecoder) request(trackIDs []string) (*http.Request, error) {
	body, err := d.body(trackIDs)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, d.host.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header = authenticationHeader(d.passphrase)
	return req, nil
}

func (d *trackDecoder) body(trackIDs []string) (io.Reader, error) {
	b, err := json.Marshal(trackIDs)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (d *trackDecoder) unmarshal(resp *http.Response) (dest []*track.Info, err error) {
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return dest, json.Unmarshal(b, &dest)
}
