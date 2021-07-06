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

	"github.com/lukasl-dev/waterlink/usecase/routeplanner/unmarkaddress"
)

const pathUnmarkAddress = "/routeplanner/free/address"

type addressUnmarker struct {
	client     *http.Client
	host       url.URL
	passphrase string
}

var _ unmarkaddress.AddressUnmarker = (*addressUnmarker)(nil)

func NewAddressUnmarker(client *http.Client, host url.URL, passphrase string) unmarkaddress.AddressUnmarker {
	host.Path += pathUnmarkAddress
	return &addressUnmarker{
		client:     client,
		host:       host,
		passphrase: passphrase,
	}
}

func (u *addressUnmarker) UnmarkAddress(address string) error {
	req, err := u.request(address)
	if err != nil {
		return err
	}
	_, err = u.client.Do(req)
	return err
}

func (u *addressUnmarker) request(address string) (*http.Request, error) {
	body, err := u.body(address)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(http.MethodPost, u.host.String(), body)
}

func (u *addressUnmarker) body(address string) (io.Reader, error) {
	type body struct {
		Address string `json:"address,omitempty"`
	}
	b, err := json.Marshal(body{Address: address})
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
