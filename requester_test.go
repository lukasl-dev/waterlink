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
	"os"
	"testing"

	"github.com/lukasl-dev/waterlink/driver/httpdriver"
	"github.com/stretchr/testify/assert"
)

func TestNewRequester(t *testing.T) {
	var (
		http           = os.Getenv("LAVALINK_HTTP")
		authentication = os.Getenv("LAVALINK_AUTHENTICATION")
	)
	host, err := url.Parse(http)
	if err != nil {
		t.Fatal(err)
	}
	opts := NewRequesterOptions().WithPassphrase(authentication)
	actual := NewRequester(*host, opts)
	expected := &requester{
		TrackDecoder:      httpdriver.NewTrackDecoder(opts.client, *host, opts.passphrase),
		TracksLoader:      httpdriver.NewTrackLoader(opts.client, *host, opts.passphrase),
		StatusGetter:      httpdriver.NewStatusGetter(opts.client, *host, opts.passphrase),
		AddressUnmarker:   httpdriver.NewAddressUnmarker(opts.client, *host, opts.passphrase),
		AddressesUnmarker: httpdriver.NewAddressesUnmarker(opts.client, *host, opts.passphrase),
	}
	assert.Equal(t, expected, actual)
}
