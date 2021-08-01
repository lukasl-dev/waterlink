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
	"github.com/lukasl-dev/waterlink/adapter/receiveevent"
	"github.com/lukasl-dev/waterlink/usecase/configureresuming"
	"github.com/lukasl-dev/waterlink/usecase/destroy"
	"github.com/lukasl-dev/waterlink/usecase/equalize"
	"github.com/lukasl-dev/waterlink/usecase/pause"
	"github.com/lukasl-dev/waterlink/usecase/play"
	"github.com/lukasl-dev/waterlink/usecase/seek"
	"github.com/lukasl-dev/waterlink/usecase/stop"
	"github.com/lukasl-dev/waterlink/usecase/updatevoice"
	"github.com/lukasl-dev/waterlink/usecase/updatevolume"
)

// Connection wraps all connection-oriented use cases.
type Connection interface {
	receiveevent.EventReceiver
	configureresuming.ResumingConfigurer
	destroy.Destroyer
	equalize.Equalizer
	pause.Pauser
	play.Player
	seek.Seeker
	stop.Stopper
	updatevoice.VoiceUpdater
	updatevolume.VolumeUpdater

	// Resumed returns true whenever the Connection has
	// been resumed.
	Resumed() bool
}

// connection is the default implementation of Connection.
type connection struct {
	receiveevent.EventReceiver
	configureresuming.ResumingConfigurer
	destroy.Destroyer
	equalize.Equalizer
	pause.Pauser
	play.Player
	seek.Seeker
	stop.Stopper
	updatevoice.VoiceUpdater
	updatevolume.VolumeUpdater

	// resumed is true whenever the Connection has
	// been resumed.
	resumed bool
}

var _ Connection = (*connection)(nil)

// Resumed returns true whenever the Connection has
// been resumed.
func (conn *connection) Resumed() bool {
	return conn.resumed
}
