package snowflake

import (
	"fmt"
	"strconv"
)

// Snowflake is the way of Discord to identify objects uniquely.
//
// See https://discord.com/developers/docs/reference#snowflakes for more information.
type Snowflake uint32

// Zero is the zero value of a Snowflake.
const Zero Snowflake = 0

// expectedLength is the expected length of characters of a snowflake.
const expectedLength = 19

var (
	// ErrParseInvalidLength occurs when the length of the snowflake mismatches the expected length.
	ErrParseInvalidLength error = &pkgError{fmt.Errorf("invalid length (must be %d)", expectedLength)}
)

// Parse parses a string into a Snowflake. It returns an error if the string provided is malformed.
func Parse(s string) (Snowflake, error) {
	if len(s) != expectedLength {
		return Zero, ErrParseInvalidLength
	}

	atoi, err := strconv.Atoi(s)
	if err != nil {
		return Zero, &pkgError{fmt.Errorf("%w", err)}
	}

	return Snowflake(atoi), nil
}

// MustParse parses a string into a Snowflake. It panics if the string provided is malformed.
func MustParse(s string) Snowflake {
	sf, err := Parse(s)
	if err != nil {
		panic(err)
	}

	return sf
}
