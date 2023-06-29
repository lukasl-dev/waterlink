package player

import (
	"encoding/json"
	"fmt"
	"github.com/lukasl-dev/waterlink/v3/lavalink/player/filter"
)

// Filters is a map of filter names to their respective filter values.
//
// By default, the following filters are available:
//   - volume: float32
//   - equalizer: []filter.EqualizerBand
//   - karaoke: *filter.Karaoke
//   - timescale: *filter.Timescale
//   - tremolo: *filter.Tremolo
//   - vibrato: *filter.Vibrato
//   - rotation: *filter.Rotation
//   - distortion: *filter.Distortion
//   - channelMix: *filter.ChannelMix
//   - lowPass: *filter.LowPass
//
// Custom filters can be added by using Lavalink plugins. Therefore, the filters
// are not limited to the ones listed above.
type Filters map[string]json.RawMessage

var (
	// ErrFilterNotFound occurs when a filter is not found in the Filters map.
	ErrFilterNotFound = fmt.Errorf("waterlink: player: filter not found")
)

// GetFilter returns the filter value for the given key. If the filter is not
// found, ErrFilterNotFound is returned.
func GetFilter[V any](f Filters, key string) (V, error) {
	var u V

	data, ok := f[key]
	if !ok {
		return u, ErrFilterNotFound
	}

	if err := json.Unmarshal(data, &u); err != nil {
		return u, fmt.Errorf("waterlink: player: failed to unmarshal filter: %w", err)
	}

	return u, nil
}

// Volume returns the volume filter value in decimal form. Range is 0.0 to 5.0,
// where 1.0 is 100%. Values above 1.0 may cause clipping.
func (f Filters) Volume() (float32, error) {
	return GetFilter[float32](f, "volume")
}

// Equalizer returns the equalizer bands used by the player. It allows
// adjustments of 15 different bands.
func (f Filters) Equalizer() ([]filter.EqualizerBand, error) {
	return GetFilter[[]filter.EqualizerBand](f, "equalizer")
}

// Karaoke returns the karaoke filter used by the player. It allows elimination
// parts of a band, usually targeting vocals.
func (f Filters) Karaoke() (*filter.Karaoke, error) {
	return GetFilter[*filter.Karaoke](f, "karaoke")
}

// Timescale returns the timescale filter used by the player. It allows change
// of speed, pitch and rate.
func (f Filters) Timescale() (*filter.Timescale, error) {
	return GetFilter[*filter.Timescale](f, "timescale")
}

// Tremolo returns the tremolo filter used by the player. It allows creation of
// a shuddering effect, where the volume quickly oscillates.
func (f Filters) Tremolo() (*filter.Tremolo, error) {
	return GetFilter[*filter.Tremolo](f, "tremolo")
}

// Vibrato returns the vibrato filter used by the player. It allows creation of
// a shuddering effect, where the pitch quickly oscillates.
func (f Filters) Vibrato() (*filter.Vibrato, error) {
	return GetFilter[*filter.Vibrato](f, "vibrato")
}

// Rotation returns the rotation filter used by the player. It allows rotation of
// sound around the stereo channels/user headphones, a.k.a. Audio Panning.
func (f Filters) Rotation() (*filter.Rotation, error) {
	return GetFilter[*filter.Rotation](f, "rotation")
}

// Distortion returns the distortion filter used by the player. It allows creation
// of a distortion effect on the audio.
func (f Filters) Distortion() (*filter.Distortion, error) {
	return GetFilter[*filter.Distortion](f, "distortion")
}

// ChannelMix returns the channel mix filter used by the player. It allows
// mixing of the channels of the audio.
func (f Filters) ChannelMix() (*filter.ChannelMix, error) {
	return GetFilter[*filter.ChannelMix](f, "channelMix")
}

// LowPass returns the low pass filter used by the player. It allows filtering
// of higher frequencies.
func (f Filters) LowPass() (*filter.LowPass, error) {
	return GetFilter[*filter.LowPass](f, "lowPass")
}
