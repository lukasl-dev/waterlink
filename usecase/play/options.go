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

type Options struct {
	StartTime uint `json:"startTime,omitempty"`
	EndTime   uint `json:"endTime,omitempty"`
	Volume    int  `json:"volume,omitempty"`
	NoReplace bool `json:"noReplace,omitempty"`
	Paused    bool `json:"paused,omitempty"`
}

func NewOptions() *Options {
	return new(Options)
}

func (opts *Options) WithStartTime(startTime uint) *Options {
	opts.StartTime = startTime
	return opts
}

func (opts *Options) WithEndTime(endTime uint) *Options {
	opts.EndTime = endTime
	return opts
}

func (opts *Options) WithVolume(volume int) *Options {
	opts.Volume = volume
	return opts
}

func (opts *Options) WithNoReplace(noReplace bool) *Options {
	opts.NoReplace = noReplace
	return opts
}

func (opts *Options) WithPaused(paused bool) *Options {
	opts.Paused = paused
	return opts
}
