package discord

import (
	"fmt"
	"strconv"
)

// Snowflake is the way of Discord to identify objects uniquely.
//
// See https://discord.com/developers/docs/reference#snowflakes for more information.
type Snowflake uint32

// ParseSnowflake parses a string into a Snowflake. It returns an error if the string provided is malformed.
func ParseSnowflake(s string) (Snowflake, error) {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("waterlink: discord: %w", err)
	}

	return Snowflake(atoi), nil
}

// MustParseSnowflake parses a string into a Snowflake. It panics if the string provided is malformed.
func MustParseSnowflake(s string) Snowflake {
	sf, err := ParseSnowflake(s)
	if err != nil {
		panic(err)
	}

	return sf
}

// String formats the Snowflake into a string.
func (sf Snowflake) String() string {
	return strconv.Itoa(int(sf))
}
