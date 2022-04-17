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

package track

// Info holds the information about a Track.
type Info struct {
	// Identifier is the audio source specific identifier.
	Identifier string `json:"identifier,omitempty"`

	// Seekable indicates whether seeking is available.
	Seekable bool `json:"isSeekable,omitempty"`

	// Author is the name of the author of the Info.
	Author string `json:"author,omitempty"`

	// Length corresponds to the duration of the Track
	// in milliseconds.
	Length uint `json:"length,omitempty"`

	// Stream indicates if the Track is a stream.
	Stream bool `json:"isStream,omitempty"`

	// Position indicates the current position of the Track.
	Position uint `json:"position,omitempty"`

	// Title is the name of the Track.
	Title string `json:"title,omitempty"`

	// URI is the URL or the local path to the audio source.
	URI string `json:"uri,omitempty"`

	// SourceName is the name of the audio source.
	SourceName string `json:"sourceName,omitempty"`
}
