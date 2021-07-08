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

import "net/http"

// RequesterOptions is used to configure further specifications
// of the Requester.
type RequesterOptions struct {
	client     *http.Client
	passphrase string
}

// NewRequesterOptions returns a new RequesterOptions.
func NewRequesterOptions() *RequesterOptions {
	return &RequesterOptions{
		client: http.DefaultClient,
	}
}

// minimizeRequesterOptions minimizes the passed options
// to a single one.
func minimizeRequesterOptions(opts []*RequesterOptions) *RequesterOptions {
	if len(opts) > 0 {
		return opts[0]
	}
	return NewRequesterOptions()
}

// WithClient sets the client to the parameter value.
func (opts *RequesterOptions) WithClient(client *http.Client) *RequesterOptions {
	opts.client = client
	return opts
}

// WithPassphrase sets the passphrase to the parameter value.
func (opts *RequesterOptions) WithPassphrase(passphrase string) *RequesterOptions {
	opts.passphrase = passphrase
	return opts
}
