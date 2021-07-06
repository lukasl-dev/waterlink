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

	"github.com/lukasl-dev/waterlink/entity/routeplanner"
	"github.com/lukasl-dev/waterlink/usecase/routeplanner/getstatus"
)

const pathGetStatus = "/routeplanner/status"

type statusGetter struct {
	client     *http.Client
	host       url.URL
	passphrase string
}

var _ getstatus.StatusGetter = (*statusGetter)(nil)

func NewStatusGetter(client *http.Client, host url.URL, passphrase string) getstatus.StatusGetter {
	host.Path += pathGetStatus
	return &statusGetter{
		client:     client,
		host:       host,
		passphrase: passphrase,
	}
}

func (g *statusGetter) Status() (*routeplanner.Status, error) {
	req, err := g.request()
	if err != nil {
		return nil, err
	}
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	return g.unmarshal(resp)
}

func (g *statusGetter) request() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, g.host.String(), nil)
}

func (g *statusGetter) unmarshal(resp *http.Response) (dest *routeplanner.Status, err error) {
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return dest, json.Unmarshal(b, &dest)
}
