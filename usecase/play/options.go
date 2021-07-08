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

package play

// Options is used to configure further specifications
// of the Player.Play method.
type Options struct {
	// StartTime defines the time from which the
	// track is to be played.
	StartTime uint `json:"startTime,omitempty"`

	// EndTime defines the time until the track is
	// to be played.
	EndTime uint `json:"endTime,omitempty"`

	// Volume defines the volume at which the track
	// is played.
	Volume uint `json:"volume,omitempty"`

	// NoReplace defines whether this action should be
	// ignored if another track is currently playing.
	NoReplace bool `json:"noReplace,omitempty"`

	// Paused defines whether the playback should be paused.
	Paused bool `json:"paused,omitempty"`
}

// NewOptions returns a new Options.
func NewOptions() *Options {
	return new(Options)
}

// MinimizeOptions minimizes the passed options to a
// single one.
func MinimizeOptions(opts ...*Options) *Options {
	if len(opts) > 0 {
		return opts[0]
	}
	return NewOptions()
}

// WithStartTime sets the start time to the parameter
// value.
func (opts *Options) WithStartTime(startTime uint) *Options {
	opts.StartTime = startTime
	return opts
}

// WithEndTime sets the end time to the parameter value.
func (opts *Options) WithEndTime(endTime uint) *Options {
	opts.EndTime = endTime
	return opts
}

// WithVolume sets the volume to the parameter value.
func (opts *Options) WithVolume(volume uint) *Options {
	opts.Volume = volume
	return opts
}

// WithNoReplace sets no replace to the parameter value.
func (opts *Options) WithNoReplace(noReplace bool) *Options {
	opts.NoReplace = noReplace
	return opts
}

// WithPaused sets paused to the parameter value.
func (opts *Options) WithPaused(paused bool) *Options {
	opts.Paused = paused
	return opts
}
