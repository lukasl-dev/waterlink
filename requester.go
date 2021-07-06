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
	"net/url"

	"github.com/lukasl-dev/waterlink/driver/httpdriver"
	"github.com/lukasl-dev/waterlink/usecase/decodetrack"
	"github.com/lukasl-dev/waterlink/usecase/loadtrack"
	"github.com/lukasl-dev/waterlink/usecase/routeplanner/getstatus"
	"github.com/lukasl-dev/waterlink/usecase/routeplanner/unmarkaddress"
	"github.com/lukasl-dev/waterlink/usecase/routeplanner/unmarkaddresses"
)

type Requester interface {
	decodetrack.TrackDecoder
	loadtrack.TrackLoader
	getstatus.StatusGetter
	unmarkaddress.AddressUnmarker
	unmarkaddresses.AddressesUnmarker
}

type requester struct {
	decodetrack.TrackDecoder
	loadtrack.TrackLoader
	getstatus.StatusGetter
	unmarkaddress.AddressUnmarker
	unmarkaddresses.AddressesUnmarker
}

var _ Requester = (*requester)(nil)

func NewRequester(host url.URL, opts ...*RequesterOptions) Requester {
	m := minimizeRequesterOptions(opts)
	return &requester{
		TrackDecoder:      httpdriver.NewTrackDecoder(m.client, host, m.passphrase),
		TrackLoader:       httpdriver.NewTrackLoader(m.client, host, m.passphrase),
		StatusGetter:      httpdriver.NewStatusGetter(m.client, host, m.passphrase),
		AddressUnmarker:   httpdriver.NewAddressUnmarker(m.client, host, m.passphrase),
		AddressesUnmarker: httpdriver.NewAddressesUnmarker(m.client, host, m.passphrase),
	}
}