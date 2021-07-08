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

package websocketdriver

import (
	"github.com/gorilla/websocket"
	"github.com/lukasl-dev/waterlink/usecase/configureresuming"
)

type resumeConfigurer struct {
	conn *websocket.Conn
}

var _ configureresuming.ResumingConfigurer = (*resumeConfigurer)(nil)

func NewResumeConfigurer(conn *websocket.Conn) configureresuming.ResumingConfigurer {
	return &resumeConfigurer{
		conn: conn,
	}
}

type configureResumingPayload struct {
	OP      op     `json:"op,omitempty"`
	Key     string `json:"key,omitempty"`
	Timeout uint   `json:"timeout,omitempty"`
}

func (c *resumeConfigurer) ConfigureResuming(resumeKey string, timeout uint) error {
	return c.conn.WriteJSON(configureResumingPayload{
		OP:      opConfigureResuming,
		Key:     resumeKey,
		Timeout: timeout,
	})
}
