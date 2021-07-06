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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	actual := NewOptions()
	expected := new(Options)
	assert.Equal(t, expected, actual)
}

func TestMinimizeOptions(t *testing.T) {
	var (
		expected = NewOptions().WithEndTime(10)
		other    = NewOptions().WithVolume(100)
		opts     = []*Options{expected, other}
	)
	actual := MinimizeOptions(opts...)
	assert.Equal(t, expected, actual)
}

func TestOptions_WithStartTime(t *testing.T) {
	const (
		startTime = 5
	)
	actual := NewOptions().WithStartTime(startTime)
	expected := &Options{StartTime: startTime}
	assert.Equal(t, expected, actual)
}

func TestOptions_WithEndTime(t *testing.T) {
	const (
		endTime = 5
	)
	actual := NewOptions().WithEndTime(endTime)
	expected := &Options{EndTime: endTime}
	assert.Equal(t, expected, actual)
}

func TestOptions_WithVolume(t *testing.T) {
	const (
		volume = 5
	)
	actual := NewOptions().WithVolume(volume)
	expected := &Options{Volume: volume}
	assert.Equal(t, expected, actual)
}

func TestOptions_WithNoReplace(t *testing.T) {
	const (
		noReplace = true
	)
	actual := NewOptions().WithNoReplace(noReplace)
	expected := &Options{NoReplace: noReplace}
	assert.Equal(t, expected, actual)
}

func TestOptions_WithPaused(t *testing.T) {
	const (
		paused = true
	)
	actual := NewOptions().WithPaused(paused)
	expected := &Options{Paused: paused}
	assert.Equal(t, expected, actual)
}
