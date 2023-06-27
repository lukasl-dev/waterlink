package snowflake

import "fmt"

// pkgError is a wrapper for snowflake-related errors.
type pkgError struct {
	// err is the actual error occurred.
	err error
}

// Error returns the error message. All snowflake-related errors are prefixed with "snowflake: ".
func (e *pkgError) Error() string {
	return fmt.Sprintf("snowflake: %s", e.err)
}
